package account

type UpdateDaoProposalStateRequest struct {
	DAOAddress   string `json:"daoAddress"`
	ProposalType string `json:"proposalType"`
	Status       int32  `json:"status"`
}

type AllowProposalCreateReq struct {
	DAOAddress  string `json:"daoAddress"`
	SubCategory string `json:"subcategory"`
}

type QueryDAOProposalStateRequest struct {
	DAOAddress []string `json:"daoAddress"`
	// ProposalType string `json:"proposalType"`
}

type QueryProposalDecisionRequest struct {
	DaoAddress  string `json:"daoAddress"`
	SubCategory string `json:"subCategory"`
	// ProposalType string `json:"proposalType"`
}

type ProposalInfo struct {
	ProposalType string `json:"proposalType"`
	Status       int32  `json:"status"`
}

type DAOProposalInfo struct {
	ProposalInfos []ProposalInfo `json:"proposalInfo"`
}

type PositiveDecisionProposal struct {
	ProposalID string `json:"proposalID"`
	Height     int64  `json:"height"`
}

type Result struct {
	Success int64 `json:"success"`
}

type QueryAuditReportItem struct {
	DaoAddress string `json:"daoAddress"`
}

type QueryByDAOAddressReq struct {
	DAOAddress string `json:"daoAddress"`
	Status     int32  `json:"status"`
}

type ProposalReq struct {
	ProposalID string `json:"proposalID"`
}

type ProposalListInfo struct {
	DAOAddress    string `json:"daoAddress"`
	ProposalID    string `json:"proposalID"`
	Category      string `json:"category"`
	Subcategory   string `json:"subCategory"`
	AuditPeriod   int32  `json:"auditPeriod"`
	BlockHeight   int64  `json:"blockHeight"`
	Agree         int    `json:"agree"`
	ProposalTitle string `json:"proposalTitle"`
}
