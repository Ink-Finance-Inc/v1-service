package account

type SocialInfo struct {
	Addr         string `json:"addr"`
	MediaType    string `json:"mediaType"`
	MediaAccount string `json:"mediaAccount"`
}

type QuerySocialInfoRequest struct {
	Keywords []string `json:"keywords"`
}

type QuerySocialInfoResponse struct {
	Code       int64        `json:"id"`
	Message    string       `json:"message"`
	SocialInfo []SocialInfo `json:"socialInfo"`
}

type AutocompleteRequest struct {
	Keywords string `json:"keywords"`
}

type AutocompleteResponse struct {
}

type ValidSocialContentRequest struct {
	MainAccount  string
	ChainType    string
	ChainID      int32
	MediaAccount string
	MediaType    string
	Link         string
}

type ValidSocialContentResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Sign    string `json:"sign"`
}
