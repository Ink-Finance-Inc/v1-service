package services

import (
	"context"
	"sort"

	conf "inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
	request "inkfinance.xyz/request/dao"
)

type IDAOService interface {
	GetDAOList(ctx context.Context, req *request.GetDAOReq) ([]request.DAOInfo, error)
	GetDAOInfo(ctx context.Context, req *request.DAOInfoBindReq) (request.DAOInfo, error)
	BindDAOInfo(ctx context.Context, req *request.DAOInfoBindReq) (int32, error)
}

type DAOService struct{}

func NewDAOService() IDAOService {
	var service = &DAOService{
		// middleware
	}
	return service
}

func (DAOService) GetDAOInfo(ctx context.Context, req *request.DAOInfoBindReq) (request.DAOInfo, error) {
	db := query.Use(conf.DB)
	tb := db.DaoInfo
	dao := tb.WithContext(ctx)

	info, err := dao.Where(tb.DaoAddress.Eq(req.DAOAddress)).First()
	if err == nil {
		i := request.DAOInfo{}
		i.DAOLogo = info.DaoLogo
		return i, err
	}
	return request.DAOInfo{}, err

}

func (DAOService) BindDAOInfo(ctx context.Context, req *request.DAOInfoBindReq) (int32, error) {
	db := query.Use(conf.DB)
	tb := db.DaoInfo
	dao := tb.WithContext(ctx)

	info, err := dao.Where(tb.DaoAddress.Eq(req.DAOAddress)).First()
	if err != nil {
		info = &model.DaoInfo{}
		info.DaoAddress = req.DAOAddress
		info.DaoLogo = req.DAOLogo
	} else {
		info.DaoLogo = req.DAOLogo
	}

	dao.Save(info)
	return 1, nil

}

func (DAOService) GetDAOList(ctx context.Context, req *request.GetDAOReq) ([]request.DAOInfo, error) {
	db := query.Use(conf.DB)
	tb := db.DaoList
	dao := db.DaoList.WithContext(ctx)
	tbSocial := db.SocialInfo
	tbDAOInfo := db.DaoInfo
	var results []request.DAOInfo
	dao.LeftJoin(tbSocial, tbSocial.MainAccounts.EqCol(tb.AdminWallet)).LeftJoin(tbDAOInfo, tbDAOInfo.DaoAddress.EqCol(tb.DaoAddress)).Select(tbDAOInfo.DaoLogo.As("DAOLogo"), tbSocial.SocialAccount.As("MediaAccount"), tbSocial.MediaType.As("MediaType"), tb.DaoAddress.As("DAOAddress"), tb.AdminWallet.As("AdminWallet"), tb.BlockHeight.As("BlockHeight")).ScanByPage(&results, int((req.Page-1)*req.Size), int(req.Size))

	for i := 0; i < len(results); i++ {
		if results[i].MediaAccount != "" {
			results[i].AdminFullInfo = results[i].MediaAccount + "@" + results[i].MediaType + "<" + results[i].AdminWallet + ">"
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].BlockHeight >= results[j].BlockHeight
	})

	var daos []string

	for i := 0; i < len(results); i++ {
		daos = append(daos, results[i].DAOAddress)
	}

	// add dao proposals
	var daoCount []request.DAOProposalCount
	tbProposal := db.ProposalDecision

	err := tbProposal.WithContext(ctx).Select(tbProposal.DaoAddress.As("DAOAddress"), tbProposal.ProposalID.Count().As("Count")).Where(tbProposal.DaoAddress.In(daos...)).Group(tbProposal.DaoAddress).Scan(&daoCount)

	if err == nil {

		for i := 0; i < len(results); i++ {
			results[i].ProposalCount = getProposlCount(results[i].DAOAddress, daoCount)
		}

	}

	return results, nil
}

func getProposlCount(dao string, daoCount []request.DAOProposalCount) int32 {
	for i := 0; i < len(daoCount); i++ {
		if daoCount[i].DAOAddress == dao {
			return daoCount[i].Count
		}
	}
	return 0
}
