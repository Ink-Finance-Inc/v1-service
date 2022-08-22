package bang_verse

import (
	"context"
	"net/http"
	"reflect"

	request "inkfinance.xyz/request/proposal"
	services "inkfinance.xyz/service"
	transport "inkfinance.xyz/transport/base"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func MakeUpdateProposalStatus(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.UpdateProposalStatus)),
		transport.BuildDecodeRequest(&request.UpdateDaoProposalStateRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeQueryDAOInfoService(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.QueryDAOInfo)),
		transport.BuildDecodeRequest(&request.QueryDAOProposalStateRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeQueryProposalDecisionService(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.QueryProposalDecisionInfo)),
		transport.BuildDecodeRequest(&request.QueryProposalDecisionRequest{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeQueryProposalDecisionCountService(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetDAOProposalDecisionCount)),
		transport.BuildDecodeRequest(&request.DAOProposalReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
func MakeProposalListHandler(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetProposalByDAO)),
		transport.BuildDecodeRequest(&request.QueryByDAOAddressReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeAllowToCreateProposalHandler(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.AllowToCreate)),
		transport.BuildDecodeRequest(&request.AllowProposalCreateReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeProposalMessageHandler(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetProposalMessage)),
		transport.BuildDecodeRequest(&request.ProposalReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}

func MakeProposalStatusHandler(svc services.ProposalService) http.Handler {
	handler := httpTransport.NewServer(
		transport.BuildEndpoint(reflect.ValueOf(svc.GetProposalStatus)),
		transport.BuildDecodeRequest(&request.ProposalReq{}),
		transport.EncodeResponse,
		transport.ErrorServerOption(),
		httpTransport.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
			return ctx
		}),
	)
	return handler
}
