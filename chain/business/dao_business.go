package business

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/chain/FactoryManager"
	"inkfinance.xyz/chain/MasterDAO"
	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

type NewDAOHandler struct {
}

// event IncomeReport(address indexed dao, address indexed operator, bytes data);
func (r NewDAOHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	db := query.Use(conf.DB)
	contractEventDAO := db.ContractEventsProgress.WithContext(context.Background())
	daoListDAO := db.DaoList.WithContext(context.Background())
	contract := req.(*FactoryManager.FactoryManager)
	iter, err := contract.FilterECreateContract(opts, nil, nil, nil)

	if err == nil {
		// currentBlockHeight, err := client.BlockNumber(context.Background())
		if err != nil {
		}

		// Address, err := abi.NewType("address", "", nil)
		// if err != nil {
		// 	// return nil, abi.New(err, true)
		// }
		// addressArg := abi.Arguments{
		// 	{
		// 		Type: abi.AddressTy,
		// 	},
		// }
		Bytes32, err := abi.NewType("bytes32", "", nil)
		if err != nil {
			// return nil, abi.New(err, true)
		}
		arg := abi.Arguments{
			{
				Type: Bytes32,
			},
		}
		/*
			event ECreateContract(
				address indexed admin,
				bytes32 indexed reqID,
				address indexed deployAddr,
				bytes32 factoryType,
				bytes32 factoryID
			);
		*/
		// abi, err := abi.JSON(strings.NewReader(FactoryManager.FactoryManagerMetaData.ABI))
		// return perr.New(err, true)
		for iter.Next() {
			targetEvent := iter.Event

			fmt.Println("RAW:\n", targetEvent.Raw)
			fmt.Println("RAW DATA:\n", targetEvent.Raw.Data)

			p0 := targetEvent.Raw.Topics[0].String()
			p1 := common.HexToAddress(targetEvent.Raw.Topics[1].String())
			p2 := targetEvent.Raw.Topics[2].String()
			p3 := common.HexToAddress(targetEvent.Raw.Topics[3].String())
			h := targetEvent.Raw.BlockNumber
			// ct, err3 := addressArg.Unpack(targetEvent.DeployAddr[:])

			// if type = 0xec5e7c7d22597af3c657ad4890c44990eedd38c4ded579e8861cf76cfaf77710;
			s1, err1 := arg.Unpack(targetEvent.FactoryType[:])
			s2, err2 := arg.Unpack(targetEvent.FactoryID[:])
			// s1, err1 := abi.Constructor.Inputs.Unpack(targetEvent.FactoryType[:])

			fmt.Println(err1, ",,,,", err2)
			fmt.Println("add: ", p3, "\np0: ", p0, "\nadmin: ", p1, "\nreqID: ", p2, "\ndeployAddr: ", p3, "\nfactoryType: ", s1, "\nfactoryID: ", s2, "\ns1[0]: ")
			// crypto.Keccak256Hash("")
			if len(s1) > 0 {
				factoryType := string("0x" + hex.EncodeToString(targetEvent.FactoryType[:]))
				// fmt.Println("compare :::---:::\n", string("0x"+hex.EncodeToString(targetEvent.FactoryType[:])), "\n", crypto.Keccak256Hash([]byte("addresstag.FACTORY_DAO")))
				if factoryType == "0xec5e7c7d22597af3c657ad4890c44990eedd38c4ded579e8861cf76cfaf77710" {
					contractEventDAO.Create(&model.ContractEventsProgress{Events: "UCV_CREATE_MONITOR", ContractAddress: p3.String(), StartHeight: 0, CurrentHeight: int64(h), Valid: 1, UpdateTime: time.Now()})
					daoListDAO.Create(&model.DaoList{DaoAddress: p3.String(), AdminWallet: p1.String(), BlockHeight: int64(h), BlockTime: time.Now().Unix()})
				}
			}

		}
	}

	return nil
}

type BindUCVHandler struct {
}

// event IncomeReport(address indexed dao, address indexed operator, bytes data);
func (r BindUCVHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {

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

	return nil
}
