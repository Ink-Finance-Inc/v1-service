package chain

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"

	"inkfinance.xyz/chain/AuditReporter"
	"inkfinance.xyz/chain/FactoryManager"
	"inkfinance.xyz/chain/InkPayManage"
	"inkfinance.xyz/chain/MasterDAO"
	"inkfinance.xyz/chain/ProposalExecMessageBoard"
	"inkfinance.xyz/chain/ProposalRegistry"
	chainbiz "inkfinance.xyz/chain/business"
)

/**

	// 已经写过的报告 -
	// Proposal + Duration + (KvData) 明确字段 + 5/31 / 30/ ，  6/15， 7/15， 8/15 ，Duration, StartTime
	// 判断 有数据的 报表。 第几期 -
	DAO ID +
	AuditReport -
		审计报告内容
		时间 + 时间
	发起提案 -
		解析成功提案时间
		提安内容Duration
		提案第一次的时间，生效时间
	// duration 去掉*30
**/

var client *ethclient.Client

// var client
func StartChainScan() {
	checkStatus()
}

func initClient() error {

	_client, err := ethclient.Dial(viper.GetString("block_service"))

	if err == nil {
		client = _client
		return nil
	} else {
		fmt.Println(err)
		return err
	}
}

func checkStatus() {
	err := initClient()

	if err == nil {
		go scan()
	} else {
		eventScan := time.NewTimer(time.Second * 2)
		for {
			select {
			case <-eventScan.C:
				err = initClient()
				if err == nil {
					go scan()
				} else {
					eventScan.Reset(time.Second * 2)
				}
			}
		}
	}
}

func scan() {
	eventScan := time.NewTimer(time.Second * 2)

	for {
		select {
		case <-eventScan.C:
			println("event scan timer")

			db := query.Use(conf.DB)
			dao := db.ContractEventsProgress.WithContext(context.Background())
			progress, err := dao.Where(db.ContractEventsProgress.Valid.Eq(1)).Find()

			if err == nil {
				for i := 0; i < len(progress); i++ {
					log.Println("handle:", progress[i])
					handleContractEventProgress(progress[i])
				}
			}
			eventScan.Reset(time.Second * 2)
		}
	}
}

func handleContractEventProgress(contractEvent *model.ContractEventsProgress) {
	contractAddress := common.HexToAddress(contractEvent.ContractAddress)
	if contractEvent.Events == ("ProposalDecision") {
		contractInterface, err := ProposalRegistry.NewProposalRegistry(contractAddress, client)
		if err == nil {
			var handler1 = chainbiz.ProposalRegistryDecisionHandler{}
			var handler2 = chainbiz.ProposalRegistryNewProposalHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler1, handler2)
		}

	} else if contractEvent.Events == "DAO_CREATE_MONITOR" {
		// Monitor Factory Create Action, then start to monitor

		contractInterface, err := FactoryManager.NewFactoryManager(contractAddress, client)
		if err == nil {
			var handler1 = chainbiz.NewDAOHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler1)
		}

		// handleNewDaoMonitor(contractEvent)
	} else if contractEvent.Events == "UCV_CREATE_MONITOR" {
		// Create

		contractInterface, err := MasterDAO.NewMasterDAO(contractAddress, client)
		if err == nil {
			var handler1 = chainbiz.BindUCVHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler1)
		}

		// handleUCVBindMonitor(contractEvent)

	} else if contractEvent.Events == "UCV_MONITOR" {
		// Monitor UCV Event
		// handlePayrollBusiness(contractEvent)
		contractInterface, err := InkPayManage.NewInkPayManage(contractAddress, client)
		if err == nil {
			var handler1 = chainbiz.PayrollSignHandler{}
			var handler2 = chainbiz.PayrollSetPayScheduleHandler{}
			var handler3 = chainbiz.PayrollWithdrawHandler{}
			var handler4 = chainbiz.PayrollAddMemberHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler1, handler2, handler3, handler4)
			// handleContractProgress(client, contractEvent, contractInterface, handler3)
		}
	} else if contractEvent.Events == "INCOME_REPORT" {
		contractInterface, err := AuditReporter.NewAuditReporter(contractAddress, client)
		if err == nil {
			var handler = chainbiz.AuditReportHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler)
		}

	} else if contractEvent.Events == "OFFCHAIN_MESSAGE" {
		contractInterface, err := ProposalExecMessageBoard.NewProposalExecMessageBoard(contractAddress, client)
		if err == nil {
			var handler = chainbiz.MessageHandler{}
			handleContractProgress(client, contractEvent, contractInterface, handler)
		}

	}
}

func handleContractProgress(client *ethclient.Client, contractEvent *model.ContractEventsProgress, contractInterface interface{}, handler ...chainbiz.IEventHandler) error {
	fmt.Println("to handle: ", handler)
	db := query.Use(conf.DB)

	currentBlockHeight, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("bErr:", err)
		return err
	}
	fmt.Println("block height is :", currentBlockHeight)

	currentScanHeight := uint64(contractEvent.CurrentHeight)

	if currentBlockHeight <= currentScanHeight {
		fmt.Println("nothing to scan")
		return nil
	}

	if currentScanHeight <= uint64(contractEvent.StartHeight) {
		currentScanHeight = uint64(contractEvent.StartHeight)
	}

	scanFrom, scanTo := calcStartEndBlock(currentBlockHeight, currentScanHeight)

	for {
		if scanFrom >= scanTo {
			// record
			dao := db.ContractEventsProgress.WithContext(context.Background())
			contractEvent.UpdateTime = time.Now()
			contractEvent.CurrentHeight = int64(scanTo)
			saveErr := dao.Save(contractEvent)
			if saveErr != nil {
				fmt.Println("save process error")
			}
			break
		}

		opts := &bind.FilterOpts{
			Start:   scanFrom,
			End:     &scanTo,
			Context: context.Background(),
		}

		for _, h := range handler {
			fmt.Println("handler to handle -- ", contractInterface)
			err := h.HandleEvent(client, opts, contractInterface, contractEvent)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("now:", scanFrom, scanTo)
		scanFrom, scanTo = calcStartEndBlock(currentBlockHeight, scanTo)
		fmt.Println("next:", scanFrom, scanTo)
	}

	return nil
}

func calcStartEndBlock(blockHeight uint64, current uint64) (scanFrom uint64, scanTo uint64) {
	reliableHeight := blockHeight - 10
	echoTimeScan := 1000

	if current+1 < reliableHeight {
		scanFrom = current + 1
		if current+uint64(echoTimeScan) >= reliableHeight {
			scanTo = reliableHeight
		} else {
			scanTo = current + uint64(echoTimeScan)
		}
		return scanFrom, scanTo
	} else {
		return current, current
	}
}

/*
func handleUCVBindMonitor(contractEvent *model.ContractEventsProgress) {
	db := query.Use(conf.DB)

	currentBlockHeight, bErr := client.BlockNumber(context.Background())
	if bErr != nil {
		fmt.Println("bErr:", bErr)
		return
	}
	fmt.Println("block height is :", currentBlockHeight)

	currentScanHeight := uint64(contractEvent.CurrentHeight)
	if bErr != nil {
		fmt.Println(bErr)
	}

	if currentBlockHeight <= currentScanHeight {
		fmt.Println("nothing to scan")
		return
	}

	if currentScanHeight <= uint64(contractEvent.StartHeight) {
		currentScanHeight = uint64(contractEvent.StartHeight)
	}

	scanFrom, scanTo := calcStartEndBlock(currentBlockHeight, currentScanHeight)
	contractAddress := common.HexToAddress(contractEvent.ContractAddress)

	for {
		if scanFrom >= scanTo {
			// record
			dao := db.ContractEventsProgress.WithContext(context.Background())
			contractEvent.UpdateTime = time.Now()
			contractEvent.CurrentHeight = int64(scanTo)
			saveErr := dao.Save(contractEvent)
			if saveErr != nil {
				fmt.Println("save process error")
			}
			break
		}

		opts := &bind.FilterOpts{
			Start:   scanFrom,
			End:     &scanTo,
			Context: context.Background(),
		}

		contract, err := MasterDAO.NewMasterDAO(contractAddress, client)
		if err != nil {
			fmt.Println("new contract err", err)
			return
		}

		handleUCVBind(opts, contract, contractEvent)

		fmt.Println("now:", scanFrom, scanTo)
		scanFrom, scanTo = calcStartEndBlock(currentBlockHeight, scanTo)
		fmt.Println("next:", scanFrom, scanTo)
	}
}

// event EAddUCV(address indexed ucv);
func handleUCVBind(opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) {

	contract := req.(*MasterDAO.MasterDAO)
	// currentBlockHeight, err := client.BlockNumber(context.Background())
	db := query.Use(conf.DB)
	contractEventDAO := db.ContractEventsProgress.WithContext(context.Background())
	daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())

	iter, err := contract.FilterEAddUCV(opts, nil)
	if err == nil {
		// return perr.New(err, true)
		for iter.Next() {

			targetEvent := iter.Event
			// ucv contract monitor
			h := targetEvent.Raw.BlockNumber
			ucvAddr := common.HexToAddress(targetEvent.Raw.Topics[1].String())

			fmt.Println("handle binding ..... --------------------------------------", ucvAddr.String(), parent.ContractAddress)
			contractEventDAO.Create(&model.ContractEventsProgress{Events: "UCV_MONITOR", ContractAddress: ucvAddr.String(), StartHeight: 0, CurrentHeight: int64(h), Valid: 1, UpdateTime: time.Now()})
			daoUcvMappingDAO.Create(&model.DaoUcvMapping{DaoAddress: parent.ContractAddress, UcvAddress: ucvAddr.String(), CreateTime: time.Now(), Valid: 1})
		}

		fmt.Println("handleUCVBind -----  over ")
	}
}*/

// func handleNewDaoMonitor(contractEvent *model.ContractEventsProgress) {

// 	db := query.Use(conf.DB)

// 	currentBlockHeight, bErr := client.BlockNumber(context.Background())
// 	if bErr != nil {
// 		fmt.Println("bErr:", bErr)
// 		return
// 	}
// 	fmt.Println("block height is :", currentBlockHeight)

// 	currentScanHeight := uint64(contractEvent.CurrentHeight)
// 	if bErr != nil {
// 		fmt.Println(bErr)
// 	}

// 	if currentBlockHeight <= currentScanHeight {
// 		fmt.Println("nothing to scan")
// 		return
// 	}

// 	if currentScanHeight <= uint64(contractEvent.StartHeight) {
// 		currentScanHeight = uint64(contractEvent.StartHeight)
// 	}

// 	scanFrom, scanTo := calcStartEndBlock(currentBlockHeight, currentScanHeight)
// 	contractAddress := common.HexToAddress(contractEvent.ContractAddress)

// 	for {
// 		if scanFrom >= scanTo {
// 			// record
// 			dao := db.ContractEventsProgress.WithContext(context.Background())
// 			contractEvent.UpdateTime = time.Now()
// 			contractEvent.CurrentHeight = int64(scanTo)
// 			saveErr := dao.Save(contractEvent)
// 			if saveErr != nil {
// 				fmt.Println("save process error")
// 			}
// 			break
// 		}

// 		opts := &bind.FilterOpts{
// 			Start:   scanFrom,
// 			End:     &scanTo,
// 			Context: context.Background(),
// 		}

// 		contract, err := FactoryManager.NewFactoryManager(contractAddress, client)
// 		if err != nil {
// 			fmt.Println("new contract err", err)
// 			return
// 		}

// 		handleDAOCreate(opts, contract, contractEvent)

// 		fmt.Println("now:", scanFrom, scanTo)
// 		scanFrom, scanTo = calcStartEndBlock(currentBlockHeight, scanTo)
// 		fmt.Println("next:", scanFrom, scanTo)
// 	}
// }

// func handlePayrollBusiness(contractEvent *model.ContractEventsProgress) {

// 	db := query.Use(conf.DB)
// 	// ProposalregistryABI

// 	currentBlockHeight, bErr := client.BlockNumber(context.Background())
// 	if bErr != nil {
// 		fmt.Println("bErr:", bErr)
// 		return
// 	}
// 	fmt.Println("block height is :", currentBlockHeight)
// 	currentScanHeight := uint64(contractEvent.CurrentHeight)
// 	if bErr != nil {
// 		fmt.Println(bErr)
// 	}

// 	if currentBlockHeight <= currentScanHeight {
// 		fmt.Println("nothing to scan")
// 		return
// 	}

// 	if currentScanHeight <= uint64(contractEvent.StartHeight) {
// 		currentScanHeight = uint64(contractEvent.StartHeight)
// 	}

// 	scanFrom, scanTo := calcStartEndBlock(currentBlockHeight, currentScanHeight)
// 	contractAddress := common.HexToAddress(contractEvent.ContractAddress)

// 	for {
// 		if scanFrom >= scanTo {
// 			// record
// 			dao := db.ContractEventsProgress.WithContext(context.Background())
// 			contractEvent.UpdateTime = time.Now()
// 			contractEvent.CurrentHeight = int64(scanTo)
// 			saveErr := dao.Save(contractEvent)
// 			if saveErr != nil {
// 				fmt.Println("save process error")
// 			} else {

// 				fmt.Println("save process done from ", scanFrom, scanTo)
// 			}
// 			break
// 		}

// 		opts := &bind.FilterOpts{
// 			Start:   scanFrom,
// 			End:     &scanTo,
// 			Context: context.Background(),
// 		}

// 		contract, err := InkPayManage.NewInkPayManage(contractAddress, client)
// 		if err != nil {
// 			fmt.Println("new contract err", err)
// 			return
// 		}

// 		// contractEvent.ContractAddress

// 		fmt.Println("start to handle sign ------------------------")
// 		// handleSign(opts, contract, contractEvent)
// 		// handleAddPayMember(opts, contract, contractEvent)
// 		// handleSetPaySchedule(opts, contract, contractEvent)
// 		// handleWithdraw(opts, contract, contractEvent)
// 		fmt.Println("done to handleSetPaySchedule  ------------------------")
// 		fmt.Println("now:", scanFrom, scanTo)
// 		scanFrom, scanTo = calcStartEndBlock(currentBlockHeight, scanTo)
// 		fmt.Println("next:", scanFrom, scanTo)
// 	}

// }

/**

 */
// func handleDAOCreate(opts *bind.FilterOpts, contract *FactoryManager.FactoryManager, parent *model.ContractEventsProgress) {
// 	db := query.Use(conf.DB)
// 	contractEventDAO := db.ContractEventsProgress.WithContext(context.Background())
// 	iter, err := contract.FilterECreateContract(opts, nil, nil, nil)

// 	if err == nil {
// 		// currentBlockHeight, err := client.BlockNumber(context.Background())
// 		if err != nil {
// 		}

// 		// Address, err := abi.NewType("address", "", nil)
// 		// if err != nil {
// 		// 	// return nil, abi.New(err, true)
// 		// }
// 		// addressArg := abi.Arguments{
// 		// 	{
// 		// 		Type: abi.AddressTy,
// 		// 	},
// 		// }
// 		Bytes32, err := abi.NewType("bytes32", "", nil)
// 		if err != nil {
// 			// return nil, abi.New(err, true)
// 		}
// 		arg := abi.Arguments{
// 			{
// 				Type: Bytes32,
// 			},
// 		}
// 		/*
// 			event ECreateContract(
// 				address indexed admin,
// 				bytes32 indexed reqID,
// 				address indexed deployAddr,
// 				bytes32 factoryType,
// 				bytes32 factoryID
// 			);
// 		*/
// 		// abi, err := abi.JSON(strings.NewReader(FactoryManager.FactoryManagerMetaData.ABI))
// 		// return perr.New(err, true)
// 		for iter.Next() {
// 			targetEvent := iter.Event

// 			fmt.Println("RAW:\n", targetEvent.Raw)
// 			fmt.Println("RAW DATA:\n", targetEvent.Raw.Data)

// 			p0 := targetEvent.Raw.Topics[0].String()
// 			p1 := targetEvent.Raw.Topics[1].String()
// 			p2 := targetEvent.Raw.Topics[2].String()
// 			p3 := common.HexToAddress(targetEvent.Raw.Topics[3].String())
// 			h := targetEvent.Raw.BlockNumber
// 			// ct, err3 := addressArg.Unpack(targetEvent.DeployAddr[:])

// 			// if type = 0xec5e7c7d22597af3c657ad4890c44990eedd38c4ded579e8861cf76cfaf77710;
// 			s1, err1 := arg.Unpack(targetEvent.FactoryType[:])
// 			s2, err2 := arg.Unpack(targetEvent.FactoryID[:])
// 			// s1, err1 := abi.Constructor.Inputs.Unpack(targetEvent.FactoryType[:])

// 			fmt.Println(err1, ",,,,", err2)
// 			fmt.Println("add: ", p3, "\np0: ", p0, "\nadmin: ", p1, "\nreqID: ", p2, "\ndeployAddr: ", p3, "\nfactoryType: ", s1, "\nfactoryID: ", s2, "\ns1[0]: ")
// 			// crypto.Keccak256Hash("")
// 			if len(s1) > 0 {
// 				factoryType := string("0x" + hex.EncodeToString(targetEvent.FactoryType[:]))
// 				// fmt.Println("compare :::---:::\n", string("0x"+hex.EncodeToString(targetEvent.FactoryType[:])), "\n", crypto.Keccak256Hash([]byte("addresstag.FACTORY_DAO")))
// 				if factoryType == "0xec5e7c7d22597af3c657ad4890c44990eedd38c4ded579e8861cf76cfaf77710" {
// 					contractEventDAO.Create(&model.ContractEventsProgress{Events: "UCV_CREATE_MONITOR", ContractAddress: p3.String(), StartHeight: 0, CurrentHeight: int64(h), Valid: 1, UpdateTime: time.Now()})
// 				}
// 			}

// 		}
// 	}
// }

// // event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin);
// func handleWithdraw(opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) {
// 	fmt.Println("handle handleWithdraw  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// 	contract := req.(*InkPayManage.InkPayManage)
// 	db := query.Use(conf.DB)
// 	scheduleWithdrawLogDAO := db.ScheduleWithdrawLog.WithContext(context.Background())

// 	iter, err := contract.FilterWithdraw(opts, nil, nil, nil)
// 	if err == nil {

// 		if err != nil {
// 			return
// 		}
// 		// return perr.New(err, true)
// 		for iter.Next() {

// 			targetEvent := iter.Event

// 			scheduleID := targetEvent.Raw.Topics[1] //arg.Unpack(targetEvent.Raw.Topics[1])
// 			// addr := targetEvent.Raw.Topics[2]
// 			// times := targetEvent.Raw.Topics[3]
// 			// total := targetEvent.Total
// 			tokenAddress := targetEvent.Coin
// 			// select by dao / schedule / member first, if exist update
// 			scheduleWithdrawLogDAO.Create(&model.ScheduleWithdrawLog{ScheduleID: int32(scheduleID.Big().Int64()), MemberAddress: "", TokenAddress: tokenAddress.Hex(), UpdateTim: time.Now(), CreateTime: time.Now()})

// 		}
// 	}
// 	fmt.Println("handle handleSetPaySchedule finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// }

/**
  event ESetPaySchedule(uint256 indexed id, uint256 duration, uint256 times, uint256 startTime);
*/
// func handleSetPaySchedule(opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) {
// 	fmt.Println("handle handleSetPaySchedule  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// 	contract := req.(*InkPayManage.InkPayManage)
// 	db := query.Use(conf.DB)
// 	scheduleInfoDAO := db.ScheduleInfo.WithContext(context.Background())

// 	iter, err := contract.FilterESetPaySchedule(opts, nil)
// 	if err == nil {
// 		daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
// 		daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()
// 		if err != nil {
// 			return
// 		}
// 		// return perr.New(err, true)
// 		for iter.Next() {

// 			targetEvent := iter.Event

// 			scheduleID := targetEvent.Raw.Topics[1] //arg.Unpack(targetEvent.Raw.Topics[1])
// 			duration := targetEvent.Duration
// 			times := targetEvent.Times
// 			startTime := targetEvent.StartTime

// 			scheduleInfoDAO.Create(&model.ScheduleInfo{ScheduleID: int32(scheduleID.Big().Int64()), Duration: duration.Int64(), ScheduleTimes: times.Int64(), StartTime: startTime.Int64(), CreateTime: time.Now(), DaoAddress: daoAddress.DaoAddress})

// 		}
// 	}
// 	fmt.Println("handle handleSetPaySchedule finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// }

/*

  event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc);

  /*

 	bytes desc is SetPayMember

 	struct SetPayMember {
    address addr; // media account
    address coin;
    uint256 oncePay;
    bytes desc;   // Title@@Description
  }


  event ERemoveMember(uint256 indexed id, address indexed addr);
  event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin);
*/

//   event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc);
// func handleAddPayMember(opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) {
// 	fmt.Println("handle handleAddPayMember  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// 	contract := req.(*InkPayManage.InkPayManage)
// 	db := query.Use(conf.DB)
// 	scheduleMemberDAO := db.ScheduleMemberInfo.WithContext(context.Background())

// 	iter, err := contract.FilterEAddMember(opts, nil, nil)

// 	if err == nil {
// 		daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
// 		daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()
// 		if err != nil {
// 			return
// 		}

// 		stringTy, err := abi.NewType("string", "", nil)
// 		if err != nil {
// 			// return nil, abi.New(err, true)
// 		}
// 		arg := abi.Arguments{
// 			{
// 				Type: stringTy,
// 			},
// 		}
// 		// return perr.New(err, true)
// 		for iter.Next() {
// 			// event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc);
// 			targetEvent := iter.Event

// 			scheduleID := targetEvent.Raw.Topics[1]
// 			memberAddr := common.HexToAddress(targetEvent.Raw.Topics[2].String())
// 			tokenAddr := targetEvent.Coin
// 			oncePay := targetEvent.OncePay
// 			titleAndDesc, err := arg.Unpack(targetEvent.Desc)

// 			titleAndDescStr := string(titleAndDesc[0].(string))

// 			fmt.Println("*** *** *** *** *** ****** *** ****** *** ****** *** ****** *** ***::::", titleAndDescStr, err)
// 			// fmt.Println("scheduleID:", scheduleID, "\ntimeID:", timeID, "\nsigner:", signer)
// 			contentSlice := strings.Split(titleAndDescStr, "@@")

// 			title := contentSlice[0]
// 			desc := ""
// 			if len(contentSlice) >= 2 {
// 				desc = contentSlice[1]
// 			}
// 			fmt.Println(targetEvent.Desc, "  DESC::: *** *** *** *** *** ****** *** ****** *** ****** *** ****** *** *** : ", desc)

// 			scheduleMemberDAO.Create(&model.ScheduleMemberInfo{DaoAddress: daoAddress.DaoAddress, ScheduleID: int32(scheduleID.Big().Int64()), MemberAddr: memberAddr.String(), TokenAddr: tokenAddr.String(), OncePay: oncePay.Int64(), ListTitle: title, Description: desc, CreateTime: time.Now(), UpdateTime: time.Now(), Valid: 1})

// 		}
// 	}
// 	fmt.Println("handle handleAddPayMember finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// }

/**
  event Sign(uint256 indexed id, uint256 indexed timeID, address indexed addr);
  // record signed schedule_signed_info
*/

// func handleSign(opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) {
// 	fmt.Println("handle sign ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// 	contract := req.(*InkPayManage.InkPayManage)
// 	db := query.Use(conf.DB)
// 	scheduleSignDAO := db.ScheduleSignedInfo.WithContext(context.Background())
// 	daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
// 	daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()

// 	iter, err := contract.FilterSign(opts, nil, nil, nil)
// 	if err == nil {
// 		// return perr.New(err, true)
// 		for iter.Next() {
// 			targetEvent := iter.Event

// 			// p0 := targetEvent.Raw.Topics[0].String()
// 			scheduleID := targetEvent.Raw.Topics[1]
// 			timeID := targetEvent.Raw.Topics[2].String()
// 			signer := common.HexToAddress(targetEvent.Raw.Topics[3].String())
// 			fmt.Println("scheduleID:", scheduleID, "\ntimeID:", timeID, "\nsigner:", signer)
// 			scheduleSignDAO.Create(&model.ScheduleSignedInfo{ScheduleID: int32(scheduleID.Big().Int64()), TimeID: int32(targetEvent.TimeID.Int64()), SignAddress: signer.String(), DaoAddress: daoAddress.DaoAddress, CreateTime: time.Now()})
// 		}
// 	}

// 	fmt.Println("handle sign finished ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
// }

/*

func handleProposalLog(contractEvent *model.ContractEventsProgress) {

	db := query.Use(conf.DB)
	// ProposalregistryABI
	contractAddress := common.HexToAddress(contractEvent.ContractAddress)
	contract, err := NewChain(contractAddress, client)
	if err != nil {
		fmt.Println("new contract err", err)
		return
	}

	currentBlockHeight, bErr := client.BlockNumber(context.Background())
	if bErr != nil {
		fmt.Println("bErr:", bErr)
		return
	}
	fmt.Println("block height is :", currentBlockHeight)

	currentScanHeight := uint64(contractEvent.CurrentHeight)
	if bErr != nil {
		fmt.Println(bErr)
	}

	if currentBlockHeight <= currentScanHeight {
		fmt.Println("nothing to scan")
		return
	}

	if currentScanHeight <= uint64(contractEvent.StartHeight) {
		currentScanHeight = uint64(contractEvent.StartHeight)
	}

	scanFrom, scanTo := calcStartEndBlock(currentBlockHeight, currentScanHeight)

	for {
		if scanFrom >= scanTo {
			// record

			dao := db.ContractEventsProgress.WithContext(context.Background())
			contractEvent.UpdateTime = time.Now()
			contractEvent.CurrentHeight = int64(scanTo)
			saveErr := dao.Save(contractEvent)
			if saveErr != nil {
				fmt.Println("save process error")
			}
			break
		}

		opts := &bind.FilterOpts{
			Start:   scanFrom,
			End:     &scanTo,
			Context: context.Background(),
		}

		log.Println("query:", opts)

		handleNewProposalRegistry(opts, contract)
		handleEProposalDecide(opts, contract)

		fmt.Println("now:", scanFrom, scanTo)
		scanFrom, scanTo = calcStartEndBlock(currentBlockHeight, scanTo)
		fmt.Println("next:", scanFrom, scanTo)
	}

}

func handleEProposalDecide(opts *bind.FilterOpts, contract *Chain) {
	db := query.Use(conf.DB)
	proposalDecisionDAO := db.ProposalDecision.WithContext(context.Background())
	iter, err := contract.FilterEProposalDecide(opts, nil, nil, nil)

	if err == nil {
		// return perr.New(err, true)
		for iter.Next() {

			targetEvent := iter.Event

			proposalID := targetEvent.Raw.Topics[2].String()
			decision, err := proposalDecisionDAO.Where(db.ProposalDecision.DaoAddress.Eq(targetEvent.Dao.String()), db.ProposalDecision.ProposalID.Eq(proposalID)).First()

			fmt.Println("find proposal decision :", proposalID, targetEvent.Dao.String())

			if err == nil {
				decision.UpdateTime = time.Now()
				decision.Agree = targetEvent.Agree
				proposalDecisionDAO.Save(decision)
			} else {
				fmt.Println("no decision error :", err, proposalID, targetEvent.Dao.String())
			}

		}
	}
}



func handleNewProposalRegistry(opts *bind.FilterOpts, contract *Chain) {

	db := query.Use(conf.DB)
	proposalDecisionDAO := db.ProposalDecision.WithContext(context.Background())
	stringTy, err := abi.NewType("string", "", nil)
	if err != nil {
		// return nil, perr.New(err, true)
	}

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		// return nil, abi.New(err, true)
	}

	bytesTy, err := abi.NewType("bytes", "", nil)
	if err != nil {
		// return nil, perr.New(err, true)
	}

	valueTy, err := abi.NewType("string", "", nil)
	if err != nil {
		// return nil, perr.New(err, true)
	}

	arg := abi.Arguments{
		{
			Type: stringTy,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: bytesTy,
		},
		{
			Type: valueTy,
		},
	}

	iter, err := contract.FilterENewProposal(opts, nil, nil)
	if err != nil {
		// return perr.New(err, true)

	}

	for iter.Next() {
		targetEvent := iter.Event
		subCategory := ""
		h := targetEvent.Raw.BlockNumber
		for i := 0; i < len(targetEvent.Metadata); i++ {

			// fmt.Println("metadata:", i)
			s, _ := arg.Unpack(targetEvent.Metadata[i])

			// fmt.Println(s)
			if i == 1 {
				arg1 := abi.Arguments{
					{
						Type: bytesTy,
					},
				}
				s2, _ := arg1.Unpack(s[2].([]byte))
				// fmt.Println("Subcategory: \n0---- ", s[0], "\n1------ ", s[1], "\n2--------", s[2], "\n3----------", s[3], "\nunpack-- ", s2, "\ndecode----- ", string(s2[0].([]byte)))
				subCategory = string(s2[0].([]byte))
				break
			}
		}

		dao := targetEvent.Dao.String()
		proposal := targetEvent.Raw.Topics[2].String()
		decision := &model.ProposalDecision{ProposalID: proposal, DaoAddress: dao, Subcategory: subCategory, UpdateTime: time.Now(), NewProposalHeight: int64(h)}
		proposalDecisionDAO.Save(decision)

	}

}
*/
