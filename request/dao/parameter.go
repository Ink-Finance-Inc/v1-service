package account

type GetDAOReq struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type DAOInfoBindReq struct {
	DAOAddress string `json:"daoAddress"`
	DAOLogo    string `json:"daoLogo"`
}

type DAOInfo struct {
	DAOAddress    string `json:"daoAddress"`
	AdminWallet   string `json:"adminWallet"`
	BlockHeight   int64  `json:"blockHeight"`
	AdminFullInfo string `json:"adminFullInfo"`
	MediaType     string
	MediaAccount  string
	ProposalCount int32  `json:"proposalCount"`
	DAOLogo       string `json:"daoLogo"`
}

type DAOProposalCount struct {
	DAOAddress string `json:"daoAddress"`
	Count      int32  `json:"count"`
}
