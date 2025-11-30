package loginService

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	login "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/login"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/infra/rpc"
	"github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *login.LoginReq) (resp *login.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	r, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		AccountNum: req.AccountNum,
		Password:   req.Password,
	})
	if err != nil {
		errMsg := err.Error()
		// 区分业务错误和系统错误
		if errMsg == "user not found" || errMsg == "wrong password" {
			resp = &login.LoginResp{
				Code: consts.StatusBadRequest, // 业务错误码: 账号或密码错误
				Msg:  fmt.Sprintf("账号或密码错误: %s", errMsg),
				Data: "",
			}
			return resp, nil
		}

		resp = &login.LoginResp{
			Code: consts.StatusInternalServerError,
			Msg:  "系统繁忙，请稍后再试",
			Data: "",
		}
		return resp, nil
	}
	resp = &login.LoginResp{
		Code: consts.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(r.UserId, 10),
	}
	return
}
