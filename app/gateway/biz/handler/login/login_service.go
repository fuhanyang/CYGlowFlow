package login

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/service/loginService"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/utils"
	login "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/login"
	commonUtils "github.com/fuhanyang/CYGlowFlow/common/utils"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req login.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := loginService.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	if resp.Code != consts.StatusOK {
		utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
		return
	}

	userID, err := strconv.ParseInt(resp.Data, 10, 64)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	token, err := commonUtils.GenerateToken(userID)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	c.Response.Header.Set("Authorization", "Bearer "+token)

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req login.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := loginService.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// TODO: 获取真实 userID
	userID := int64(123) // 临时模拟，实际应该从 service 返回

	token, err := commonUtils.GenerateToken(userID)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	c.Response.Header.Set("Authorization", "Bearer "+token)

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Logoff .
// @router /user/logoff [POST]
func Logoff(ctx context.Context, c *app.RequestContext) {
	var err error
	var req login.LogoffReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := loginService.NewLogoffService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
