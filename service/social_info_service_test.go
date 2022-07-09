package services

import (
	"context"
	"fmt"
	"testing"

	conf "inkfinance.xyz/conf"

	request "inkfinance.xyz/request/social"
)

func TestSocialInfoService(t *testing.T) {
	srv := setupSocialInfoervice()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	if srv != nil {

	}
	var keywords []string
	keywords = append(keywords, "aa")
	keywords = append(keywords, "bb")
	req := request.QuerySocialInfoRequest{Keywords: keywords}

	resp, err := srv.QueryMediaInfo(ctx, &req)

	if err != nil {
		fmt.Println(resp)
	}

}

// func setup() (srv IAccountService, ctx context.Context) {
func setupSocialInfoervice() (srv ISocialInfoService) {
	conf.DB = conf.ConnectDB(conf.MySQLDSN).Debug()
	return NewSocialInfoService()
}
