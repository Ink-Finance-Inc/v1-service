package bang_verse

import (
	"context"
	"net/http"
	"reflect"

	services "inkfinance.xyz/service"

	httpTransport "github.com/go-kit/kit/transport/http"
	request "inkfinance.xyz/request/util"
	transport "inkfinance.xyz/transport/base"
)

func MakeSetKeyHandler(svc services.UtilService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.SetKeyValue)),
		transport.BuildDecodeRequest(&request.KeyValueCache{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeGetKeyHandler(svc services.UtilService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetKeyValue)),
		transport.BuildDecodeRequest(&request.KeyValueCache{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeDataExistHandler(svc services.UtilService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.DataExist)),
		transport.BuildDecodeRequest(&request.GetDataExistReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
