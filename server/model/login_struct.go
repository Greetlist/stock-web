package model

type LoginRequest struct {
    Account string `json:"account"`
    Password string `json:"passwd"`
}

type LogoutRequest struct {
    Account string `json:"account"`
}

type LoginResponse struct {
    LoginSucc bool `json:"login_succ"`
    ErrMsg string `json:"err_msg"`
}

type LogoutResponse struct {
    LogoutSucc bool `json:"logout_succ"`
    ErrMsg string `json:"err_msg"`
}
