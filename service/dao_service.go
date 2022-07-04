package services

import (
	"context"
	"sort"

	conf "inkfinance.xyz/conf"
	"inkfinance.xyz/dal/query"
	request "inkfinance.xyz/request/dao"
)

type IDAOService interface {
	GetDAOList(ctx context.Context, req *request.GetDAOReq) ([]request.DAOInfo, error)
}

type DAOService struct{}

func NewDAOService() IDAOService {
	var service = &DAOService{
		// middleware
	}
	return service
}

func (DAOService) GetDAOList(ctx context.Context, req *request.GetDAOReq) ([]request.DAOInfo, error) {
	db := query.Use(conf.DB)
	tb := db.DaoList
	dao := db.DaoList.WithContext(ctx)
	tbSocial := db.SocialInfo
	var results []request.DAOInfo
	dao.LeftJoin(tbSocial, tbSocial.MainAccounts.EqCol(tb.AdminWallet)).Select(tbSocial.SocialAccount.As("MediaAccount"), tbSocial.MediaType.As("MediaType"), tb.DaoAddress.As("DAOAddress"), tb.AdminWallet.As("AdminWallet"), tb.BlockHeight.As("BlockHeight")).ScanByPage(&results, int((req.Page-1)*req.Size), int(req.Size))

	for i := 0; i < len(results); i++ {
		if results[i].MediaAccount != "" {
			results[i].AdminFullInfo = results[i].MediaAccount + "@" + results[i].MediaType + "<" + results[i].AdminWallet + ">"
		}

	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].BlockHeight >= results[j].BlockHeight
	})
	return results, nil
}
