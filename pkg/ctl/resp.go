package ctl

type ErrResponse struct {
	Err string
}

type DataResponse struct {
	Msg  interface{}
	Data interface{}
}

type TokenResponse struct {
	Msg  interface{}
	User interface{}
	Data interface{}
}
