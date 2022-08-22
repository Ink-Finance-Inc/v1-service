package bang_verse

import (
	"context"
	"net/http"
	"reflect"

	request "inkfinance.xyz/request/dao"
	services "inkfinance.xyz/service"
	transport "inkfinance.xyz/transport/base"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func MakeGetDAOListService(svc services.DAOService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetDAOList)),
		transport.BuildDecodeRequest(&request.GetDAOReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeGetDAOInfoService(svc services.DAOService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetDAOInfo)),
		transport.BuildDecodeRequest(&request.DAOInfoBindReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeDAOInfoBindService(svc services.DAOService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.BindDAOInfo)),
		transport.BuildDecodeRequest(&request.DAOInfoBindReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
