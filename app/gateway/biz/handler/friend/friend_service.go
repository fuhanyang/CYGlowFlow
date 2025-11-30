package friend

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/service/friendService"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/utils"
	friend "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/friend"
)

// GetFriendList .
// @router /friend/list [POST]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req friend.GetFriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &friend.GetFriendListResp{}
	resp, err = friendService.NewGetFriendListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddFriend .
// @router /friend/addition/with_account_num [POST]
func AddFriend(ctx context.Context, c *app.RequestContext) {
	var err error
	var req friend.AddFriendReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &friend.AddFriendResp{}
	resp, err = friendService.NewAddFriendService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
