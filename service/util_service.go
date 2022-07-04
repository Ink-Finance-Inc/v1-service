package services

import (
	"context"

	request "inkfinance.xyz/request/util"

	conf "inkfinance.xyz/conf"

	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
)

type IUtilService interface {
	GetKeyValue(ctx context.Context, req *request.KeyValueCache) (string, error)
	SetKeyValue(ctx context.Context, req *request.KeyValueCache) (bool, error)
	DataExist(ctx context.Context, req *request.GetDataExistReq) bool
}

type UtilService struct{}

func NewUtilService() IUtilService {
	var service = &UtilService{
		// middleware
	}
	return service
}

func (UtilService) GetKeyValue(ctx context.Context, req *request.KeyValueCache) (string, error) {

	db := query.Use(conf.DB)
	tbKeyValue := db.KeyValueCache
	keyValueDAO := tbKeyValue.WithContext(ctx)

	kv, err := keyValueDAO.Where(tbKeyValue.KeyIs.Eq(req.KeyIs)).First()
	if err != nil {
		return "", nil
	}

	return kv.ValueIs, nil

}

func (UtilService) SetKeyValue(ctx context.Context, req *request.KeyValueCache) (bool, error) {
	db := query.Use(conf.DB)
	tbKeyValue := db.KeyValueCache
	keyValueDAO := tbKeyValue.WithContext(ctx)
	kv, err := keyValueDAO.Where(tbKeyValue.KeyIs.Eq(req.KeyIs)).First()
	if err != nil {
		keyValueDAO.Create(&model.KeyValueCache{KeyIs: req.KeyIs, ValueIs: req.ValueIs})
		return true, nil
	}
	if kv.ValueIs != req.ValueIs {
		kv.ValueIs = req.ValueIs
		err = keyValueDAO.Save(kv)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return true, nil
	}

}

func (UtilService) DataExist(ctx context.Context, req *request.GetDataExistReq) bool {
	db := query.Use(conf.DB)

	if req.DataType == "DAO" {
		tb := db.DaoList
		dao := tb.WithContext(ctx)

		ob, _ := dao.Where(tb.BlockHeight.Gte(req.Height)).Count()
		if ob > 0 {
			return true
		}
		return false
	}

	if req.DataType == "PROPOSAL" {
		tb := db.ProposalDecision
		dao := tb.WithContext(ctx)

		ob, _ := dao.Where(tb.NewProposalHeight.Gte(req.Height)).Count()
		if ob > 0 {
			return true
		}
		return false
	}

	if req.DataType == "WITHDRAW" {
		tb := db.ScheduleWithdrawLog
		dao := tb.WithContext(ctx)

		ob, _ := dao.Where(tb.BlockHeight.Gte(req.Height)).Count()
		if ob > 0 {
			return true
		}
		return false
	}

	if req.DataType == "AUDITREPORT" {
		tb := db.AuditReport
		dao := tb.WithContext(ctx)

		ob, _ := dao.Where(tb.BlockHeight.Gte(req.Height)).Count()
		if ob > 0 {
			return true
		}
		return false
	}

	return false

}
