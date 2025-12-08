package workflow

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"

	"github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager/executionmanagerservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() executionmanagerservice.Client
	Service() string
	ExecuteWorkflow(ctx context.Context, Req *exeution_manager.ExecuteWorkflowRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error)
	ListExecutions(ctx context.Context, Req *exeution_manager.ExecutionQueryRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error)
	GetExecutionDetail(ctx context.Context, Req *exeution_manager.GetExecutionDetailRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error)
	CancelExecution(ctx context.Context, Req *exeution_manager.CancelExecutionRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := executionmanagerservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient executionmanagerservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() executionmanagerservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ExecuteWorkflow(ctx context.Context, Req *exeution_manager.ExecuteWorkflowRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error) {
	return c.kitexClient.ExecuteWorkflow(ctx, Req, callOptions...)
}

func (c *clientImpl) ListExecutions(ctx context.Context, Req *exeution_manager.ExecutionQueryRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error) {
	return c.kitexClient.ListExecutions(ctx, Req, callOptions...)
}

func (c *clientImpl) GetExecutionDetail(ctx context.Context, Req *exeution_manager.GetExecutionDetailRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error) {
	return c.kitexClient.GetExecutionDetail(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelExecution(ctx context.Context, Req *exeution_manager.CancelExecutionRequest, callOptions ...callopt.Option) (r *common.BaseResponse, err error) {
	return c.kitexClient.CancelExecution(ctx, Req, callOptions...)
}
