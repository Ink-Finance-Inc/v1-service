package business

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/chain/ProposalExecMessageBoard"
	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

/*
   event EMessage(
       bytes32 indexed topicID,
       address indexed addr,
       bytes32 indexed execProposalID,
       bytes32 nowProposalID,
       bytes message
   );
*/
type MessageHandler struct {
}

func (r MessageHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("start to handle MessageHandler")

	contract := req.(*ProposalExecMessageBoard.ProposalExecMessageBoard)

	db := query.Use(conf.DB)
	tb := db.ExecuteMessage
	dao := tb.WithContext(context.Background())
	iter, err := contract.FilterEMessage(opts, nil, nil, nil)

	if err == nil {

		if err != nil {
			// return nil, abi.New(err, true)
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

		for iter.Next() {
			targetEvent := iter.Event

			topicID := targetEvent.Raw.Topics[1]
			theAddress := common.HexToAddress(targetEvent.Raw.Topics[2].Hex())
			execProposalID := targetEvent.Raw.Topics[3]
			nowProposal := targetEvent.NowProposalID
			message, err := arg.Unpack(targetEvent.Message)

			if err != nil {

			}

			// nowProposalUnpacked, err := arg.Unpack(nowProposal[:])

			nowProposalID := string("0x" + hex.EncodeToString(targetEvent.NowProposalID[:]))
			fmt.Println("now proposal: ", nowProposalID)
			fmt.Println(topicID)
			fmt.Println(theAddress)
			fmt.Println(execProposalID)
			fmt.Println(string(nowProposal[:]))
			fmt.Println(message)
			// model.ExecuteMessage()

			dao.Create(&model.ExecuteMessage{TopicID: topicID.Hex(), Address: theAddress.String(), ExecProposalID: execProposalID.Hex(), NowProposalID: nowProposalID, Message: message[0].(string)})

		}
	}

	return nil
}
