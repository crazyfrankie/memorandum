package controller

import "memorandum/pkg/ctl"

func ErrorResponse(err error) *ctl.ErrResponse {
	return &ctl.ErrResponse{
		Err: err.Error(),
	}
}

func DataResponse(resp, data interface{}) *ctl.DataResponse {
	return &ctl.DataResponse{
		Msg:  resp,
		Data: data,
	}
}

func TokenResponse(msg, user, data interface{}) *ctl.TokenResponse {
	return &ctl.TokenResponse{
		Msg:  msg,
		User: user,
		Data: data,
	}
}
