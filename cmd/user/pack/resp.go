package pack

import (
	"errors"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
