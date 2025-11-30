package loginService

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	login "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/login"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/infra/rpc"
	"github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *login.RegisterReq) (resp *login.RegisterResp, err error) {
	// 1. 参数校验
	if req.Password != req.PasswordConfirm {
		resp = &login.RegisterResp{
			Code: consts.StatusBadRequest, // 业务错误码：密码不一致
			Msg:  "两次输入的密码不一致",
		}
		return resp, nil
	}

	// 2. 调用 RPC
	r, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		// Email: req.Email, // 如果 IDL 中有 Email 字段
		// Phone: req.Phone, // 如果 IDL 中有 Phone 字段
	})
	if err != nil {
		resp = &login.RegisterResp{
			Code: consts.StatusInternalServerError,
			Msg:  err.Error(), // 或者统一提示 "注册失败"
		}
		return resp, nil
	}

	// 3. 返回结果
	resp = &login.RegisterResp{
		Code: consts.StatusOK,
		Msg:  "success",
		Data: strconv.FormatInt(r.AccountNum, 10), // 返回生成的账号
	}
	return resp, nil
}
