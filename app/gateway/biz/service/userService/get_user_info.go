package userService

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/user"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/infra/rpc"
	kitexUser "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/user"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(ctx context.Context, requestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: requestContext, Context: ctx}
}

func (h *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// 将gateway的account_num转换为int64类型
	accountNum, err := strconv.ParseInt(req.AccountNum, 10, 64)
	if err != nil {
		resp = &user.GetUserInfoResp{
			Code: consts.StatusBadRequest,
			Msg:  "无效的账号格式",
		}
		return resp, nil
	}

	// 调用user服务的RPC接口
	r, err := rpc.UserClient.GetUserInfo(h.Context, &kitexUser.GetUserInfoReq{
		AccountNum: accountNum,
	})
	if err != nil {
		errMsg := err.Error()
		// 区分业务错误和系统错误
		if errMsg == "user not found" {
			resp = &user.GetUserInfoResp{
				Code: consts.StatusNotFound,
				Msg:  "用户不存在",
			}
			return resp, nil
		}

		resp = &user.GetUserInfoResp{
			Code: consts.StatusInternalServerError,
			Msg:  "系统繁忙，请稍后再试",
		}
		return resp, nil
	}

	// 构建gateway层的响应数据
	userInfo := &user.UserInfo{
		AccountNum: strconv.FormatInt(r.UserInfo.AccountNum, 10),
		Name:       r.UserInfo.Name,
		Desc:       r.UserInfo.Desc,
		Email:      r.UserInfo.Email,
		Phone:      r.UserInfo.Phone,
		AvatarUrl:  r.UserInfo.AvatarUrl,
		CreateAt:   r.UserInfo.CreatedAt,
		LastLogin:  "", // user服务没有提供最后登录时间，暂时为空
		Ip:         "", // 需要从请求中获取IP
		Tags:       []string{},
		Level:      1,
		Stats: &user.UserStats{
			FriendsCount:  0, // 需要调用friend服务获取
			WorkflowCount: 0, // 需要调用workflow服务获取
		},
	}

	resp = &user.GetUserInfoResp{
		Code: consts.StatusOK,
		Msg:  "success",
		Data: userInfo,
	}
	return
}
