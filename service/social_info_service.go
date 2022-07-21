package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	request "inkfinance.xyz/request/social"

	conf "inkfinance.xyz/conf"

	model "inkfinance.xyz/dal/model"

	"inkfinance.xyz/dal/query"
)

type ISocialInfoService interface {
	ValidSocialInfo(ctx context.Context, req *request.ValidSocialContentRequest) (*request.ValidSocialContentResponse, error)
	QueryMediaoExist(ctx context.Context, req *request.ValidSocialContentRequest) (bool, error)
	QueryMediaInfo(ctx context.Context, req *request.QuerySocialInfoRequest) (*request.QuerySocialInfoResponse, error)
	AutocompleteQuery(ctx context.Context, req *request.AutocompleteRequest) (*request.QuerySocialInfoResponse, error)
}

type SocialInfoService struct{}

func NewSocialInfoService() ISocialInfoService {
	var service = &SocialInfoService{
		// middleware
	}
	return service
}

func (SocialInfoService) QueryMediaoExist(ctx context.Context, req *request.ValidSocialContentRequest) (bool, error) {
	db := query.Use(conf.DB)

	mDo := db.SocialInfo.WithContext(ctx)
	_, err := mDo.Where(db.SocialInfo.SocialAccount.Eq(req.MediaAccount), db.SocialInfo.MediaType.Eq(req.MediaType)).First()
	if err == nil {
		// exist
		return true, nil

	}

	return false, nil
}

func (SocialInfoService) ValidSocialInfo(ctx context.Context, req *request.ValidSocialContentRequest) (*request.ValidSocialContentResponse, error) {

	//
	// queryTwitter(req.Link)
	// database

	// MediaType:    req.MediaType,
	db := query.Use(conf.DB)

	mDo := db.SocialInfo.WithContext(ctx)
	socialInfo, err := mDo.Where(db.SocialInfo.MainAccounts.Eq(req.MainAccount)).Or(db.SocialInfo.SocialAccount.Eq(req.MediaAccount)).First()
	if err == nil {
		// exist
		fmt.Println("main:", socialInfo.MainAccounts, " social: ", socialInfo.SocialAccount, " ----- exist! ")
		return &request.ValidSocialContentResponse{Code: 1, Message: "success", Sign: "sign messages"}, nil
	}

	users := model.User{
		MainAccounts: req.MainAccount,
		ChainType:    req.ChainType,
		ChainID:      req.ChainID,
		ID:           0,
		CreateTime:   time.Now(),
		Valid:        1,
	}
	do := db.User.WithContext(ctx)
	do.Create(&users)

	mySocials := model.SocialInfo{
		SocialAccount: req.MediaAccount,
		MediaType:     req.MediaType,
		UserID:        users.ID,
		MainAccounts:  req.MainAccount,
		CreateTime:    time.Now(),
		Valid:         1,
	}
	mDo.Create(&mySocials)

	return &request.ValidSocialContentResponse{Code: 1, Message: "success", Sign: "sign messages"}, nil

}

func (SocialInfoService) QueryMediaInfo(ctx context.Context, req *request.QuerySocialInfoRequest) (*request.QuerySocialInfoResponse, error) {
	db := query.Use(conf.DB)
	q := db.SocialInfo
	do := q.WithContext(ctx)
	fmt.Println(req.Keywords)
	fmt.Println(req.Keywords[0])

	splitedKeys := strings.Split(req.Keywords[0], ",")

	queryResult, err := do.Where(db.SocialInfo.MainAccounts.In(splitedKeys...), db.SocialInfo.Valid.Eq(1)).Find()
	// queryResult, err := do.Where(db.SocialInfo.Valid.Eq(1)).Find()
	if err != nil {
		return nil, err
	}

	var results []request.SocialInfo

	for i := 0; i < len(queryResult); i++ {
		s := request.SocialInfo{Addr: queryResult[i].MainAccounts, MediaType: queryResult[i].MediaType, MediaAccount: queryResult[i].SocialAccount}
		results = append(results, s)
	}

	return &request.QuerySocialInfoResponse{
		Code:       200,
		Message:    "",
		SocialInfo: results,
	}, nil
}

func (SocialInfoService) AutocompleteQuery(ctx context.Context, req *request.AutocompleteRequest) (*request.QuerySocialInfoResponse, error) {
	db := query.Use(conf.DB)
	q := db.SocialInfo
	do := q.WithContext(ctx)

	queryResult, err := do.Where(db.SocialInfo.SocialAccount.Like(req.Keywords+"%"), db.SocialInfo.Valid.Eq(1)).Find()
	if err != nil {
		return nil, err
	}

	var results []request.SocialInfo

	for i := 0; i < len(queryResult); i++ {
		s := request.SocialInfo{Addr: queryResult[i].MainAccounts, MediaType: queryResult[i].MediaType, MediaAccount: queryResult[i].SocialAccount}
		results = append(results, s)
	}

	return &request.QuerySocialInfoResponse{
		Code:       200,
		Message:    "",
		SocialInfo: results,
	}, nil
}

func queryTwitter() {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets?ids=1500736276085243907", nil)

	req.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAALiGbgEAAAAAvFU4oKHtGn86RG8obizhs9enixc%3DTD0dd2A7apOL3wiUZMeTx5bd51qzRVyWH2gqL98LRZI6suogMl")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		// handle err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}
