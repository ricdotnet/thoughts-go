package models

type Request struct {
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"Message"`
}
