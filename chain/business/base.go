package business

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/dal/model"
)

type IEventHandler interface {
	HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error
}
