package model

type Respone struct{
	StatusCode int `"json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	
}