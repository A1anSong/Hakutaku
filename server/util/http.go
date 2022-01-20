package util

type ErrResp struct {
	Status int    // 0表示正常
	Msg    string // 错误信息
}

func Resp(status int, msg string) *ErrResp {
	return &ErrResp{
		Status: 0,
		Msg:    msg,
	}
}
