package base

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"inkfinance.xyz/constant"
	errors "inkfinance.xyz/error"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	PostMethod = "POST"
	GetMethod  = "GET"

	JsonContentTyp = "application/json"
	FormContentTyp = "application/x-www-form-urlencoded"
	XmlContentTyp  = "application/xml"
)

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionID")
}

// server option
func ErrorServerOption() httptransport.ServerOption {
	return httptransport.ServerErrorEncoder(MyErrorEncoder)
}
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		terr := errors.ParseError(err.Error())
		if http.StatusText(terr.Code) != "" {
			// http错误码
			w.WriteHeader(terr.Code)
		}
		w.Header().Set("Content-Type", JsonContentTyp)
		w.Write([]byte(terr.String()))
	}
}

// check access
func checkAccess(r *http.Request, method, contentTyp string) error {
	var code int
	if r.Method != method {
		code = http.StatusMethodNotAllowed
		return errors.NewError(code, http.StatusText(code))
	}
	if contentTyp != "" && r.Header.Get("Content-Type") != contentTyp {
		code = http.StatusUnsupportedMediaType
		return errors.NewError(code, http.StatusText(code))
	}

	return nil
}

func ParamsCheckAccess(r *http.Request) error {
	return checkAccess(r, GetMethod, "")
}

func JsonCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, JsonContentTyp)
}

func FormCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, FormContentTyp)
}

func XmlCheckAccess(r *http.Request) error {
	return checkAccess(r, PostMethod, XmlContentTyp)
}

func BuildEndpoint(servceFun reflect.Value) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// defer func() {
		// 	err := recover()
		// 	switch err.(type) {
		// 	case runtime.Error:
		// 		fmt.Println("runtime error:", err)
		// 	default:
		// 		fmt.Println("error:", err)
		// 	}
		// }()

		retList := servceFun.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(request)})
		if len(retList) == 0 {
			return nil, errors.NewError(500, servceFun.String()+" error")
		} else {
			if len(retList) >= 2 && retList[1].Interface() != nil {
				return nil, retList[1].Interface().(error)
			} else {
				return retList[0].Interface(), nil
			}
		}
	}
}

func BuildDecodeRequest(req interface{}) func(context.Context, *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		defer r.Body.Close()

		if r.Method == "OPTIONS" {
			return nil, errors.NewError(3001, "duplicated")
		}

		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		if err := ParseForm(r.Form, req); err != nil {
			return nil, err
		}

		processBody := ctx.Value(constant.PROCESS_BODY)
		if processBody != nil && processBody.(bool) {
			body, err2 := ioutil.ReadAll(r.Body)
			if err2 != nil {

			}
			json.Unmarshal(body, req)
		}
		return req, nil

	}
}

type JsonResponse struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
	Msg  string      `json:"message"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	// setCors(w)
	return json.NewEncoder(w).Encode(&JsonResponse{
		Code: 1000,
		Body: response,
		Msg:  "SUCCESS",
	})
}
