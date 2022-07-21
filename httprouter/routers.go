package httprouter

import (
	services "inkfinance.xyz/service"

	dao_transport "inkfinance.xyz/transport/dao"
	payment_transport "inkfinance.xyz/transport/payment"
	proposal_transport "inkfinance.xyz/transport/proposal"
	social_transport "inkfinance.xyz/transport/socialinfo"
	util_transport "inkfinance.xyz/transport/util"

	"net/http"
	// gin-swagger middleware
)

// swagger embed files
// gin-swagger middleware

func RegisterRouter(mux *http.ServeMux) {

	mux.Handle("/social/query", social_transport.MakeQuerySocialMediaHandler(services.SocialInfoService{}))
	mux.Handle("/social/exist", social_transport.MakeQuerySocialMediaExistHandler(services.SocialInfoService{}))
	mux.Handle("/social/valid", social_transport.MakeValidSocialInfoHandler(services.SocialInfoService{}))
	mux.Handle("/social/autocomplete", social_transport.MakeAutocompleteHandler(services.SocialInfoService{}))
	mux.Handle("/proposal/flag", proposal_transport.MakeUpdateProposalStatus(services.ProposalService{}))
	mux.Handle("/proposal/allow", proposal_transport.MakeAllowToCreateProposalHandler(services.ProposalService{}))
	mux.Handle("/proposal/query", proposal_transport.MakeQueryDAOInfoService(services.ProposalService{}))
	mux.Handle("/proposal/decision", proposal_transport.MakeQueryProposalDecisionService(services.ProposalService{}))
	mux.Handle("/proposal/list", proposal_transport.MakeProposalListHandler(services.ProposalService{}))
	mux.Handle("/proposal/message", proposal_transport.MakeProposalMessageHandler(services.ProposalService{}))
	mux.Handle("/proposal/status", proposal_transport.MakeProposalStatusHandler(services.ProposalService{}))

	mux.Handle("/payment/payroll", payment_transport.MakePayrollQueryHandler(services.PaymentService{}))
	mux.Handle("/payment/schedule", payment_transport.MakeGetScheduleHandler(services.PaymentService{}))
	mux.Handle("/payment/time", payment_transport.MakeGetTimeHandler(services.PaymentService{}))
	mux.Handle("/payment/collectable", payment_transport.MakePaymentCollectableHandler(services.PaymentService{}))
	mux.Handle("/payment/sign/dao", payment_transport.MakeDAOSignInfoHandler(services.PaymentService{}))
	mux.Handle("/payment/sign/schedule", payment_transport.MakeScheduleSignInfoHandler(services.PaymentService{}))
	mux.Handle("/report/tab", payment_transport.MakeAuditReportTabHandler(services.PaymentService{}))
	mux.Handle("/report/list", payment_transport.MakeAuditReportContentHandler(services.PaymentService{}))
	mux.Handle("/payment/withdraw", payment_transport.MakeWithdrawHandler(services.PaymentService{}))

	mux.Handle("/dao/list", dao_transport.MakeGetDAOListService(services.DAOService{}))

	mux.Handle("/util/key/set", util_transport.MakeSetKeyHandler(services.UtilService{}))
	mux.Handle("/util/key/get", util_transport.MakeGetKeyHandler(services.UtilService{}))
	mux.Handle("/util/data/exist", util_transport.MakeDataExistHandler(services.UtilService{}))

}
