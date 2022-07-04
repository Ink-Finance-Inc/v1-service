package account

type GetDAOReq struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type DAOInfo struct {
	DAOAddress    string `json:"daoAddress"`
	AdminWallet   string `json:"adminWallet"`
	BlockHeight   int64  `json:"blockHeight"`
	AdminFullInfo string `json:"adminFullInfo"`
	MediaType     string
	MediaAccount  string
}
