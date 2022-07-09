package endpoint

import (
	"context"
	"fmt"

	social_req "inkfinance.xyz/request/social"

	services "inkfinance.xyz/service"

	"github.com/go-kit/kit/endpoint"
)

// make endpoint             service -> endpoint
func MakeQuerySocialMediaEndpoint(svc services.SocialInfoService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(*social_req.QuerySocialInfoRequest)
		if !ok {
			return nil, nil
		}
		resp, err := svc.QueryMediaInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// // make endpoint             service -> endpoint
func MakeValidSocialInfoEndpoint(svc services.SocialInfoService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(*social_req.ValidSocialContentRequest)

		fmt.Println("req2: %s", req)
		if !ok {
			return nil, nil
		}
		resp, err := svc.ValidSocialInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// // make endpoint             service -> endpoint
func MakeAutocompleteEndpoint(svc services.SocialInfoService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(*social_req.AutocompleteRequest)

		fmt.Println("req2: %s", req)
		if !ok {
			return nil, nil
		}
		resp, err := svc.AutocompleteQuery(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
