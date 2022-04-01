package server

import(
	. "github.com/hkaya15/PicusSecurity/Week_5_Homework/api/models/api"	
)



func NewError(code int64, message string) Error{
	return Error{Code: code, Message: message}
}