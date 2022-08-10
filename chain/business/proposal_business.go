package business

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/chain/ProposalRegistry"
	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

type ProposalRegistryDecisionHandler struct {
}

type ProposalRegistryNewProposalHandler struct {
}





func (r ProposalRegistryDecisionHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("start to handle ProposalRegistryDecisionHandler report")
	contract := req.(*ProposalRegistry.ProposalRegistry)

	db := query.Use(conf.DB)
	proposalDecisionDAO := db.ProposalDecision.WithContext(context.Background())
	iter, err := contract.FilterEProposalDecide(opts, nil, nil, nil)

	if err == nil {
		// return perr.New(err, true)
		for iter.Next() {

			targetEvent := iter.Event

			proposalID := targetEvent.Raw.Topics[2].String()
			decision, err := proposalDecisionDAO.Where(db.ProposalDecision.DaoAddress.Eq(targetEvent.Dao.String()), db.ProposalDecision.ProposalID.Eq(proposalID)).First()

			// fmt.Println(" ::::::::::::::::::::::::::::::::: find proposal decision :", proposalID, decision, targetEvent.Dao.String())

			if err == nil {

				block, err := client.BlockByHash(context.Background(), targetEvent.Raw.BlockHash)
				if err == nil {
					decision.UpdateTime = time.Now()
					if targetEvent.Agree == true {
						decision.Agree = 1
					} else {
						decision.Agree = 0
					}

					decision.AgreeTime = int64(block.Time())

					proposalDecisionDAO.Save(decision)
				} else {
					fmt.Println("err ::::::::::::::::::::::::::::::::: ", err)
				}

			} else {
				fmt.Println("no decision error :", err, proposalID, targetEvent.Dao.String())
			}

		}
	}

	return nil
}

func (r ProposalRegistryNewProposalHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("start to handle ProposalRegistryNewProposalHandler report")
	contract := req.(*ProposalRegistry.ProposalRegistry)

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
		category := ""
		proposalTitle := ""
		var period int64 = 0
		h := targetEvent.Raw.BlockNumber

		// fmt.Println("start loop kvdata :::::: ---------")
		findAuditPeriod := false
		for i := 0; i < len(targetEvent.KvData); i++ {
			s, _ := arg.Unpack(targetEvent.KvData[i])

			for j := 0; j < len(s); j++ {

				// fmt.Println(s[0].(string) == "AuditPeriod")

				if s[0].(string) == "AuditPeriod" {

					uintTy, err := abi.NewType("UintTy", "", nil)
					if err != nil {
						// return nil, perr.New(err, true)
					}
					intArg := abi.Arguments{
						{
							Type: uintTy,
						},
					}
					s1, _ := intArg.Unpack(s[2].([]byte))
					period = s1[0].(*big.Int).Int64()
					findAuditPeriod = true
					break
				}
			}

			if findAuditPeriod {
				break
			}
		}

		// fmt.Println("end loop kvdata :::::: ---------")

		// fmt.Println("length:", len(targetEvent.Metadata), targetEvent.Metadata)
		for i := 0; i < len(targetEvent.Metadata); i++ {

			s, _ := arg.Unpack(targetEvent.Metadata[i])
			fmt.Println("__________________________________________________METADATA:", s)
			if s[0].(string) == "Subcategory" {
				// fmt.Println("__________________________________________________METADATA_SUBCATEGORY: ")
				arg1 := abi.Arguments{
					{
						Type: bytesTy,
					},
				}
				s2, _ := arg1.Unpack(s[2].([]byte))
				// fmt.Println("Subcategory: \n0---- ", s[0], "\n1------ ", s[1], "\n2--------", s[2], "\n3----------", s[3], "\nunpack-- ", s2, "\ndecode----- ", string(s2[0].([]byte)))
				subCategory = string(s2[0].([]byte))

			}

			if s[0].(string) == "Category" {
				arg1 := abi.Arguments{
					{
						Type: bytesTy,
					},
				}
				s2, _ := arg1.Unpack(s[2].([]byte))
				// fmt.Println("Subcategory: \n0---- ", s[0], "\n1------ ", s[1], "\n2--------", s[2], "\n3----------", s[3], "\nunpack-- ", s2, "\ndecode----- ", string(s2[0].([]byte)))
				category = string(s2[0].([]byte))
			}

			if s[0].(string) == "ProposalTitle" {
				arg1 := abi.Arguments{
					{
						Type: bytesTy,
					},
				}
				s2, _ := arg1.Unpack(s[2].([]byte))
				// fmt.Println("Subcategory: \n0---- ", s[0], "\n1------ ", s[1], "\n2--------", s[2], "\n3----------", s[3], "\nunpack-- ", s2, "\ndecode----- ", string(s2[0].([]byte)))
				proposalTitle = string(s2[0].([]byte))
			}
		}

		dao := targetEvent.Dao.String()
		proposal := targetEvent.Raw.Topics[2].String()
		decision := &model.ProposalDecision{ProposalID: proposal, DaoAddress: dao, Subcategory: subCategory, UpdateTime: time.Now(), NewProposalHeight: int64(h), AuditPeriod: int32(period), Category: category, Agree: -1, ProposalTitle: proposalTitle}
		proposalDecisionDAO.Save(decision)

	}

	return nil
}
