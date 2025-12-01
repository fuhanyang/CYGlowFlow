package chat

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/service/chatService"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/utils"
	chat "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/chat"
)

// RefreshText .
// @router /chat/text/refresh [POST]
func RefreshText(ctx context.Context, c *app.RequestContext) {
	var err error
	var req chat.RefreshTextReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	var resp *chat.RefreshTextResp
	resp, err = chatService.NewRefreshTextService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
