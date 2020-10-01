package lmexec

import (
	"context"
	"net/http"

	lmjaeger "github.com/logicmonitor/k8s-argus/pkg/jaeger"
	"github.com/logicmonitor/k8s-argus/pkg/lmctx"
	"github.com/logicmonitor/k8s-argus/pkg/types"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/vkumbhar94/lm-sdk-go/client/lm"
	"github.com/vkumbhar94/lm-sdk-go/models"
)

// LMExec Provides utility function for SDK calls using Base object
// LMExec is holding device related api calls at the moment to mitigate rate limit handling
type LMExec struct {
	*types.Base
}

// InjectTraceComponents
func TraceClient(lctx *lmctx.LMContext) (*http.Client, context.Context) {
	span := lmjaeger.Span(lctx)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	client := &http.Client{Transport: &nethttp.Transport{}}
	return client, ctx

}

// AddDevice Add new device
func (lmexec *LMExec) AddDevice(params *lm.AddDeviceParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.AddDevice(params)
	}
}

// UpdateDevice Add new device
func (lmexec *LMExec) UpdateDevice(params *lm.UpdateDeviceParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.UpdateDevice(params)
	}
}

// GetDeviceList Add new device
func (lmexec *LMExec) GetDeviceList(params *lm.GetDeviceListParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.GetDeviceList(params)
	}
}

// PatchDevice Add new device
func (lmexec *LMExec) PatchDevice(params *lm.PatchDeviceParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.PatchDevice(params)
	}
}

// DeleteDeviceByID Add new device
func (lmexec *LMExec) DeleteDeviceByID(params *lm.DeleteDeviceByIDParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.DeleteDeviceByID(params)
	}
}

// GetImmediateDeviceListByDeviceGroupID Add new device
func (lmexec *LMExec) GetImmediateDeviceListByDeviceGroupID(params *lm.GetImmediateDeviceListByDeviceGroupIDParams) types.ExecRequest {
	return func(lctx *lmctx.LMContext) (interface{}, error) {
		client, ctx := TraceClient(lctx)
		params.SetHTTPClient(client)
		params.SetContext(ctx)
		return lmexec.LMClient.LM.GetImmediateDeviceListByDeviceGroupID(params)
	}
}

// AddDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) AddDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.AddDeviceDefault).Payload
}

// UpdateDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) UpdateDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.UpdateDeviceDefault).Payload
}

// GetDeviceListErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) GetDeviceListErrResp(err error) *models.ErrorResponse {
	return err.(*lm.GetDeviceListDefault).Payload
}

// PatchDeviceErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) PatchDeviceErrResp(err error) *models.ErrorResponse {
	return err.(*lm.PatchDeviceDefault).Payload
}

// DeleteDeviceByIDErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) DeleteDeviceByIDErrResp(err error) *models.ErrorResponse {
	return err.(*lm.DeleteDeviceByIDDefault).Payload
}

// GetImmediateDeviceListByDeviceGroupIDErrResp parse error object and returns models.ErrorResponse
func (lmexec *LMExec) GetImmediateDeviceListByDeviceGroupIDErrResp(err error) *models.ErrorResponse {
	return err.(*lm.GetImmediateDeviceListByDeviceGroupIDDefault).Payload
}
