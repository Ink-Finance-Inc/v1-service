package services

import (
	"context"
	"fmt"
	"sort"

	"gorm.io/gen"
	conf "inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
	request "inkfinance.xyz/request/proposal"
)

type IProposalService interface {
	QueryDAOInfo(ctx context.Context, req *request.QueryDAOProposalStateRequest) (*request.DAOProposalInfo, error)
	QueryProposalDecisionInfo(ctx context.Context, req *request.QueryProposalDecisionRequest) (*request.PositiveDecisionProposal, error)

	UpdateProposalStatus(ctx context.Context, req *request.UpdateDaoProposalStateRequest) (*request.Result, error)
	AllowToCreate(ctx context.Context, req *request.AllowProposalCreateReq) (int32, error)
	GetProposalByDAO(ctx context.Context, req *request.QueryByDAOAddressReq) ([]request.ProposalListInfo, error)
	GetProposalStatus(ctx context.Context, req *request.ProposalReq) (int32, error)
	GetProposalMessage(ctx context.Context, req *request.ProposalReq) (string, error)
}

type ProposalService struct{}

func NewProposalService() IProposalService {
	var service = &ProposalService{
		// middleware
	}
	return service
}

func (ProposalService) GetProposalMessage(ctx context.Context, req *request.ProposalReq) (string, error) {
	db := query.Use(conf.DB)
	tb := db.ExecuteMessage
	dao := tb.WithContext(ctx)

	m, err := dao.Where(tb.ExecProposalID.Eq(req.ProposalID)).First()
	if err != nil {
		return "", nil
	}

	return m.Message, nil

}

func (ProposalService) GetProposalStatus(ctx context.Context, req *request.ProposalReq) (int32, error) {
	db := query.Use(conf.DB)
	tb := db.ProposalDecision
	dao := db.ProposalDecision.WithContext(ctx)

	p, err := dao.Where(tb.ProposalID.Eq(req.ProposalID)).First()
	if err == nil {
		return p.Agree, nil
	}
	return -99, nil
}

func (ProposalService) AllowToCreate(ctx context.Context, req *request.AllowProposalCreateReq) (int32, error) {
	db := query.Use(conf.DB)
	tb := db.ProposalDecision
	dao := db.ProposalDecision.WithContext(ctx)

	p, err := dao.Where(tb.DaoAddress.Eq(req.DAOAddress), tb.Subcategory.Eq(req.SubCategory)).Order(tb.ID.Desc()).First()
	if err != nil {
		//no records, allow to create
		return 2, nil
	}
	return p.Agree, nil
}

/*
	DAOAddress  string `json:"daoAddress"`
	ProposalID  string `json:"propsalID"`
	Category    string `json:"category"`
	Subcategory string `json:"subCategory"`
	AuditPeriod int32  `json:"auditPeriod"`
	BlockHeight int64  `json:"blockHeight"`
	Agree       int    `json:"agree"`
*/
func (ProposalService) GetProposalByDAO(ctx context.Context, req *request.QueryByDAOAddressReq) ([]request.ProposalListInfo, error) {
	db := query.Use(conf.DB)
	tb := db.ProposalDecision
	dao := db.ProposalDecision.WithContext(ctx)

	var results []request.ProposalListInfo

	var conds []gen.Condition

	conds = append(conds, tb.DaoAddress.Eq(req.DAOAddress))
	if req.Status != 2 {
		conds = append(conds, tb.Agree.Eq(req.Status))
	}

	dao.Where(conds...).Select(tb.DaoAddress.As("DAOAddress"), tb.ProposalID.As("ProposalID"), tb.Subcategory.As("Subcategory"), tb.AuditPeriod.As("AuditPeriod"),
		tb.NewProposalHeight.As("BlockHeight"), tb.Category.As("Category"), tb.Agree.As("Agree"), tb.ProposalTitle.As("ProposalTitle")).Scan(&results)

	sort.Slice(results, func(i, j int) bool {
		return results[i].BlockHeight >= results[j].BlockHeight
	})
	return results, nil
}

func (ProposalService) QueryDAOInfo(ctx context.Context, req *request.QueryDAOProposalStateRequest) (*request.DAOProposalInfo, error) {
	db := query.Use(conf.DB)
	dao := db.ProposalStatus.WithContext(ctx)
	fmt.Println(len(req.DAOAddress))
	queryResult, err := dao.Where(db.ProposalStatus.DaoAddress.In(req.DAOAddress...)).Find()
	if err != nil {

	}

	var results []request.ProposalInfo

	for i := 0; i < len(queryResult); i++ {
		s := request.ProposalInfo{ProposalType: queryResult[i].ProposalType, Status: queryResult[i].Status}
		results = append(results, s)
	}

	return &request.DAOProposalInfo{ProposalInfos: results}, nil

}

func (ProposalService) UpdateProposalStatus(ctx context.Context, req *request.UpdateDaoProposalStateRequest) (*request.Result, error) {
	db := query.Use(conf.DB)
	dao := db.ProposalStatus.WithContext(ctx)

	queryResult, err := dao.Where(db.ProposalStatus.DaoAddress.Eq(req.DAOAddress), db.ProposalStatus.ProposalType.Eq(req.ProposalType)).First()
	if err != nil {
		fmt.Println(err)
	}

	if queryResult == nil {
		queryResult = &model.ProposalStatus{DaoAddress: req.DAOAddress, ProposalType: req.ProposalType, Status: req.Status}
		dao.Create(queryResult)
	} else {
		queryResult.Status = req.Status
		dao.Save(queryResult)
	}

	return &request.Result{Success: 1}, nil
}

func (ProposalService) QueryProposalDecisionInfo(ctx context.Context, req *request.QueryProposalDecisionRequest) (*request.PositiveDecisionProposal, error) {
	db := query.Use(conf.DB)
	dao := db.ProposalDecision.WithContext(ctx)

	queryResult, err := dao.Where(db.ProposalDecision.Agree.Eq(1), db.ProposalDecision.DaoAddress.Eq(req.DaoAddress), db.ProposalDecision.Subcategory.Eq(req.SubCategory)).Order(db.ProposalDecision.UpdateTime.Desc()).First()
	if err != nil {
		return &request.PositiveDecisionProposal{ProposalID: "", Height: 0}, nil
	} else {
		return &request.PositiveDecisionProposal{ProposalID: queryResult.ProposalID, Height: queryResult.NewProposalHeight}, nil
	}

}
