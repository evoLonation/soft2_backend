// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	UserId string `json:"user_id"`
}

type Response struct {
	Code     int    `json:"code"`
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	NickName string `json:"nick_name"`
}