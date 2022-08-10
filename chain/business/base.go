package business

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/dal/model"
)

func Try(userFn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("execute error: %v\n", err)
		}
	}()

	userFn()
}

type IEventHandler interface {
	HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error
}
