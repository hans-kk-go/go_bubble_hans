package models

// 定义请求的参数结构体

type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

// 登录请求参数

type ParamLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
