package business

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/chain/InkPayManage"
	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

type PayrollWithdrawHandler struct {
}

// event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin);
func (r PayrollWithdrawHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("handle handleWithdraw  ----- ----- ----- ----- ----- ----- ----- ----- ----- ---------------------------------------------------------------------- ")
	contract := req.(*InkPayManage.InkPayManage)
	db := query.Use(conf.DB)
	tbScheduleWithdrawLog := db.ScheduleWithdrawLog
	scheduleWithdrawLogDAO := tbScheduleWithdrawLog.WithContext(context.Background())
	daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
	daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()

	iter, err := contract.FilterWithdraw(opts, nil, nil, nil)
	if err == nil {

		if err != nil {
			return err
		}
		// uintTy, err := abi.NewType("UintTy", "", nil)
		// if err != nil {
		// 	// return nil, perr.New(err, true)
		// }
		// intArg := abi.Arguments{
		// 	{
		// 		Type: uintTy,
		// 	},
		// }
		// return perr.New(err, true)
		for iter.Next() {

			targetEvent := iter.Event
			h := int64(targetEvent.Raw.BlockNumber)
			block, err := client.BlockByHash(context.Background(), targetEvent.Raw.BlockHash)
			if err == nil {

			}
			scheduleID := targetEvent.Raw.Topics[1].Big() //arg.Unpack(targetEvent.Raw.Topics[1])
			withdrawUser := common.HexToAddress(targetEvent.Raw.Topics[2].Hex())

			hasTimeID := targetEvent.Raw.Topics[3].Big()
			total := targetEvent.Total
			tokenAddress := targetEvent.Coin

			log := model.ScheduleWithdrawLog{BlockHeight: h, DaoAddress: daoAddress.DaoAddress, ScheduleID: int32(scheduleID.Int64()), MemberAddress: withdrawUser.String(), CreateTime: int64(block.Time()), TokenAddress: tokenAddress.String(), ClaimAmount: total.String(), CurrentTimeID: int32(hasTimeID.Int64())}
			scheduleWithdrawLogDAO.Create(&log)

		}
	}
	fmt.Println("handle handleSetPaySchedule finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	return err
}

/**
  event Sign(uint256 indexed id, uint256 indexed timeID, address indexed addr);
  // record signed schedule_signed_info
*/
type PayrollSignHandler struct {
}

func (r PayrollSignHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("handle sign ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	contract := req.(*InkPayManage.InkPayManage)
	db := query.Use(conf.DB)
	scheduleSignDAO := db.ScheduleSignedInfo.WithContext(context.Background())
	daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
	daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()

	iter, err := contract.FilterSign(opts, nil, nil, nil)
	if err == nil {
		// return perr.New(err, true)
		for iter.Next() {
			targetEvent := iter.Event
			block, err := client.BlockByHash(context.Background(), targetEvent.Raw.BlockHash)
			if err == nil {

			}
			// p0 := targetEvent.Raw.Topics[0].String()
			scheduleID := targetEvent.Raw.Topics[1]
			timeID := targetEvent.Raw.Topics[2].String()
			signer := common.HexToAddress(targetEvent.Raw.Topics[3].String())
			fmt.Println("scheduleID:", scheduleID, "\ntimeID:", timeID, "\nsigner:", signer)
			scheduleSignDAO.Create(&model.ScheduleSignedInfo{ScheduleID: int32(scheduleID.Big().Int64()), TimeID: int32(targetEvent.TimeID.Int64()), SignAddress: signer.String(), DaoAddress: daoAddress.DaoAddress, CreateTime: time.Now(), BlockTime: int64(block.Time())})
		}
	}

	fmt.Println("handle sign finished ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	return err
}

//event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc);
type PayrollAddMemberHandler struct {
}

func (r PayrollAddMemberHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {

	fmt.Println("handle PayrollAddMemberHandler  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	contract := req.(*InkPayManage.InkPayManage)
	db := query.Use(conf.DB)
	scheduleMemberDAO := db.ScheduleMemberInfo.WithContext(context.Background())

	iter, err := contract.FilterEAddMember(opts, nil, nil)

	if err == nil {
		daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
		daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()
		if err != nil {
			return err
		}

		stringTy, err := abi.NewType("string", "", nil)
		if err != nil {
			// return nil, abi.New(err, true)
		}
		arg := abi.Arguments{
			{
				Type: stringTy,
			},
		}
		// return perr.New(err, true)
		for iter.Next() {
			// event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc);
			targetEvent := iter.Event

			scheduleID := targetEvent.Raw.Topics[1]
			memberAddr := common.HexToAddress(targetEvent.Raw.Topics[2].String())
			tokenAddr := targetEvent.Coin
			oncePay := targetEvent.OncePay
			titleAndDesc, err := arg.Unpack(targetEvent.Desc)
			if err == nil {

			}
			titleAndDescStr := string(titleAndDesc[0].(string))
			// fmt.Println("*** *** *** *** *** ****** *** ****** *** ****** *** ****** *** ***::::", titleAndDescStr, err)
			// fmt.Println("scheduleID:", scheduleID, "\ntimeID:", timeID, "\nsigner:", signer)
			contentSlice := strings.Split(titleAndDescStr, "@@")

			title := contentSlice[0]
			desc := ""
			if len(contentSlice) >= 2 {
				desc = contentSlice[1]
			}
			// fmt.Println(targetEvent.Desc, "  DESC::: *** *** *** *** *** ****** *** ****** *** ****** *** ****** *** *** : ", desc)
			// fmt.Println(targetEvent.Desc, "  Once PAY ::: *** *** *** *** *** ****** *** ****** *** ****** *** ****** *** *** : ", oncePay)

			scheduleMemberDAO.Create(&model.ScheduleMemberInfo{DaoAddress: daoAddress.DaoAddress, ScheduleID: int32(scheduleID.Big().Int64()), MemberAddr: memberAddr.String(), TokenAddr: tokenAddr.String(), OncePay: oncePay.String(), ListTitle: title, Description: desc, CreateTime: time.Now(), UpdateTime: time.Now(), Valid: 1})

		}
	}
	fmt.Println("handle handleAddPayMember finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	return nil
}

type PayrollSetPayScheduleHandler struct {
}

/**
  event ESetPaySchedule(uint256 indexed id, uint256 duration, uint256 times, uint256 startTime);
*/

func (r PayrollSetPayScheduleHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("handle handleSetPaySchedule  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	contract := req.(*InkPayManage.InkPayManage)
	db := query.Use(conf.DB)
	scheduleInfoDAO := db.ScheduleInfo.WithContext(context.Background())

	iter, err := contract.FilterESetPaySchedule(opts, nil)
	if err == nil {
		daoUcvMappingDAO := db.DaoUcvMapping.WithContext(context.Background())
		daoAddress, err := daoUcvMappingDAO.Where(db.DaoUcvMapping.UcvAddress.Eq(parent.ContractAddress)).First()
		if err != nil {
			return err
		}
		// return perr.New(err, true)
		for iter.Next() {

			targetEvent := iter.Event

			scheduleID := targetEvent.Raw.Topics[1] //arg.Unpack(targetEvent.Raw.Topics[1])
			duration := targetEvent.Duration
			times := targetEvent.Times
			startTime := targetEvent.StartTime

			scheduleInfoDAO.Create(&model.ScheduleInfo{ScheduleID: int32(scheduleID.Big().Int64()), Duration: duration.Int64(), ScheduleTimes: times.Int64(), StartTime: startTime.Int64(), CreateTime: time.Now(), DaoAddress: daoAddress.DaoAddress})

		}
	}
	fmt.Println("handle handleSetPaySchedule finished  ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ")
	return nil
}
