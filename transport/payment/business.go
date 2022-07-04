package bang_verse

import (
	"context"
	"net/http"
	"reflect"

	request "inkfinance.xyz/request/payment"
	services "inkfinance.xyz/service"
	transport "inkfinance.xyz/transport/base"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func MakePayrollQueryHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetPayrollList)),
		transport.BuildDecodeRequest(&request.PayrollQueryRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeGetTimeHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetTimeList)),
		transport.BuildDecodeRequest(&request.TimeQueryRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeGetScheduleHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetScheduleList)),
		transport.BuildDecodeRequest(&request.DAOScheduleRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakePaymentCollectableHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.AddressCollectable)),
		transport.BuildDecodeRequest(&request.DAOCollectableRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeDAOSignInfoHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetDAOSignInfo)),
		transport.BuildDecodeRequest(&request.DAOScheduleRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeScheduleSignInfoHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetScheduleSignInfo)),
		transport.BuildDecodeRequest(&request.TimeQueryRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeAuditReportTabHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetAuditReportTab)),
		transport.BuildDecodeRequest(&request.AuditReportRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeAuditReportContentHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetAuditReportContent)),
		transport.BuildDecodeRequest(&request.AuditReportRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeWithdrawHandler(svc services.PaymentService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetScheduleWithdraws)),
		transport.BuildDecodeRequest(&request.TimeQueryRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
