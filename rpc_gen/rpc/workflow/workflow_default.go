package workflow

import (
	"context"
	common "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/common"
	exeution_manager "github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen/workflow/infra/exeution_manager"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ExecuteWorkflow(ctx context.Context, req *exeution_manager.ExecuteWorkflowRequest, callOptions ...callopt.Option) (resp *common.BaseResponse, err error) {
	resp, err = defaultClient.ExecuteWorkflow(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ExecuteWorkflow call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListExecutions(ctx context.Context, req *exeution_manager.ExecutionQueryRequest, callOptions ...callopt.Option) (resp *common.BaseResponse, err error) {
	resp, err = defaultClient.ListExecutions(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListExecutions call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetExecutionDetail(ctx context.Context, req *exeution_manager.GetExecutionDetailRequest, callOptions ...callopt.Option) (resp *common.BaseResponse, err error) {
	resp, err = defaultClient.GetExecutionDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetExecutionDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelExecution(ctx context.Context, req *exeution_manager.CancelExecutionRequest, callOptions ...callopt.Option) (resp *common.BaseResponse, err error) {
	resp, err = defaultClient.CancelExecution(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelExecution call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
