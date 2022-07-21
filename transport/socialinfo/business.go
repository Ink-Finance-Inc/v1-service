package bang_verse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	social_request "inkfinance.xyz/request/social"
	services "inkfinance.xyz/service"

	transport "inkfinance.xyz/transport/base"

	httpTransport "github.com/go-kit/kit/transport/http"
	endpoint "inkfinance.xyz/endpoints"
)

func decodeQuerySocialRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// if err := transport.FormCheckAccess(r); err != nil {
	// 	return nil, err
	// }
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	req := &social_request.QuerySocialInfoRequest{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("r.Form:%#v\n", r.Form)
	fmt.Printf("req:%#v\n", req)
	r.Body.Close()
	return req, nil
}

func decodeValidSocialInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// if err := transport.FormCheckAccess(r); err != nil {
	// 	return nil, err
	// }
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	req := &social_request.ValidSocialContentRequest{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}

	r.Body.Close()
	return req, nil
}

func decodeAutocompleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// if err := transport.FormCheckAccess(r); err != nil {
	// 	return nil, err
	// }
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	req := &social_request.AutocompleteRequest{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}

	r.Body.Close()
	return req, nil
}

func encodeSocialResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// make handler
func MakeQuerySocialMediaHandler(svc services.SocialInfoService) http.Handler {
	handler := httpTransport.NewServer(
		endpoint.MakeQuerySocialMediaEndpoint(svc),
		decodeQuerySocialRequest,
		encodeSocialResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}

func MakeValidSocialInfoHandler(svc services.SocialInfoService) http.Handler {
	handler := httpTransport.NewServer(
		endpoint.MakeValidSocialInfoEndpoint(svc),
		decodeValidSocialInfoRequest,
		encodeSocialResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}

func MakeAutocompleteHandler(svc services.SocialInfoService) http.Handler {
	handler := httpTransport.NewServer(
		endpoint.MakeAutocompleteEndpoint(svc),
		decodeAutocompleteRequest,
		encodeSocialResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}

func MakeQuerySocialMediaExistHandler(svc services.SocialInfoService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.QueryMediaoExist)),
		transport.BuildDecodeRequest(&social_request.ValidSocialContentRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
