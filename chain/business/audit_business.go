package business

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"inkfinance.xyz/chain/AuditReporter"
	"inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

type AuditReportHandler struct {
}

// event IncomeReport(address indexed dao, address indexed operator, bytes data);
func (r AuditReportHandler) HandleEvent(client *ethclient.Client, opts *bind.FilterOpts, req interface{}, parent *model.ContractEventsProgress) error {
	fmt.Println("start to handle audit report")
	contract := req.(*AuditReporter.AuditReporter)

	db := query.Use(conf.DB)

	auditReporterDAO := db.AuditReport.WithContext(context.Background())

	iter, err := contract.FilterIncomeReport(opts, nil, nil)
	if err == nil {
		// return perr.New(err, true)
		for iter.Next() {
			targetEvent := iter.Event
			h := targetEvent.Raw.BlockNumber
			block, err := client.BlockByHash(context.Background(), targetEvent.Raw.BlockHash)
			if err == nil {
				daoAddress := common.HexToAddress(targetEvent.Raw.Topics[1].String())
				operatorAddress := common.HexToAddress(targetEvent.Raw.Topics[2].String())

				auditReporterDAO.Create(&model.AuditReport{BlockHeight: int64(h), DaoAddress: daoAddress.String(), OperatorAddress: string(operatorAddress.String()), ReportContent: targetEvent.Data, CreateTime: int64(block.Time())})
			}

		}

		fmt.Println("AuditReportHandler -----  over ")
	}

	return nil
}
