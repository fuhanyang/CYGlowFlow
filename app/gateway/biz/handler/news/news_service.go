package news

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/service/newsService"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/utils"
	news "github.com/fuhanyang/CYGlowFlow/app/gateway/hertz_gen/gateway/news"
)

// GetHotNews .
// @router /news/hot [GET]
func GetHotNews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req news.GetHotNewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &news.GetHotNewsResp{}
	resp, err = newsService.NewGetHotNewsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
