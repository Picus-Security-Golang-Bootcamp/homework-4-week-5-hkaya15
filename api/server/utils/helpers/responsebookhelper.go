package server

import (
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/models/api"
)

func NewResponseBook(code int32, message string, book *Book) APIResponseBook {
	return APIResponseBook{Code: code, Message: message, Result: book}
}

func RespondJSON(status int32, message string, payload interface{}) ResponseAny {
	return ResponseAny{Code: status, Message: message, Result: &payload}
}

type ResponseAny struct {
	Code    int32 `json:"code"`
	Message string `json:"message"`
	Result  interface{} `json:"result"`
}


