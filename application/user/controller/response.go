package controller

import "net/http"

type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

func NewResponseBody(status int, result ...any) any {
	success := status/100 == 2
	if !success {
		return nil
	}

	res := ResponseBody{
		Code:    status,
		Message: http.StatusText(status),
		Result:  nil,
	}

	switch len(result) {
	case 0:
		// skip
	case 1:
		res.Result = result[0]
	default:
		res.Result = result
	}

	return res
}
