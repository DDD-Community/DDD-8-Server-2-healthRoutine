package controller

import (
	"net/http"
)

// TODO: refactor

func ResponseByHttpStatus(status int, result ...interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	if status == http.StatusOK {
		res["code"] = status
		res["message"] = "OK"
		if len(result) <= 0 {
			return
		} else {
			res["result"] = result[0]
			return
		}
	} else if status == http.StatusCreated {
		res["code"] = status
		res["message"] = "Created"
		if len(result) <= 0 {
			return
		} else {
			res["result"] = result[0]
			return
		}
	} else {
		return nil
	}
}
