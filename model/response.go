package model

type ResponseModel struct {
	success bool
	msg     string
}

func SuccessMsg(msg string) *ResponseModel {
	return &ResponseModel{
		success: true,
		msg:     msg,
	}
}

func Success() *ResponseModel {
	return &ResponseModel{
		success: true,
		msg:     "",
	}
}
