package account

type KeyValueCache struct {
	KeyIs   string `json:"keyIs"`
	ValueIs string `json:"valueIs"`
}

type GetDataExistReq struct {
	DataType string `json:"dataType"`
	Height   int64  `json:"height"`
}
