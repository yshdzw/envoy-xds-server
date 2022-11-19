package callbacks

import (
	"context"

	discoverygrpcv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/sirupsen/logrus"
)

type callbacks struct {
	ctx           context.Context
	log           logrus.FieldLogger
	handleRequest func(request *discoverygrpcv3.DiscoveryRequest)
}

var _ serverv3.Callbacks = &callbacks{}

func NewCallbacks(ctx context.Context, log logrus.FieldLogger, handlerFunc func(request *discoverygrpcv3.DiscoveryRequest)) serverv3.Callbacks {
	return &callbacks{
		ctx:           ctx,
		log:           log,
		handleRequest: handlerFunc,
	}
}

func (c callbacks) OnFetchRequest(ctx context.Context, request *discoverygrpcv3.DiscoveryRequest) error {
	//TODO implement me
	// c.log.Debugf("OnFetchRequest..., request:%+v", request)
	c.log.Debugf("OnFetchRequest..., request:%+v", []interface{}{request.GetVersionInfo(), request.GetResponseNonce(), request.GetResourceNames(), request.GetErrorDetail()})
	return nil
}

func (c callbacks) OnFetchResponse(request *discoverygrpcv3.DiscoveryRequest, response *discoverygrpcv3.DiscoveryResponse) {
	//TODO implement me
	// c.log.Debugf("OnFetchResponse..., request:%+v, response:%+v", request, response)
	c.log.Debugf("OnFetchResponse..., request:%+v, response:%+v", []interface{}{request.GetVersionInfo(), request.GetResponseNonce(), request.GetResourceNames(), request.GetErrorDetail()}, []interface{}{response.GetVersionInfo(), response.GetNonce()})
}

func (c callbacks) OnStreamOpen(ctx context.Context, i int64, s string) error {
	//TODO implement me
	c.log.Debugf("OnStreamOpen..., i:%d, s:%s", i, s)
	return nil
}

func (c callbacks) OnStreamClosed(i int64) {
	//TODO implement me
	c.log.Debugf("OnStreamClosed..., i:%d", i)
}

func (c callbacks) OnStreamRequest(i int64, request *discoverygrpcv3.DiscoveryRequest) error {
	// c.log.Debugf("OnStreamRequest..., i:%d, request:%+v", i, request)
	c.log.Debugf("OnStreamRequest..., i:%d, request:%+v", i, []interface{}{request.GetVersionInfo(), request.GetResponseNonce(), request.GetResourceNames(), request.GetErrorDetail()})
	c.handleRequest(request)
	return nil
}

func (c callbacks) OnStreamResponse(ctx context.Context, i int64, request *discoverygrpcv3.DiscoveryRequest, response *discoverygrpcv3.DiscoveryResponse) {
	//TODO implement me
	// c.log.Debugf("OnFetchResponse..., i:%d request:%+v, response:%+v", i, request, response)
	c.log.Debugf("OnFetchResponse..., i:%d request:%+v, response:%+v", i, []interface{}{request.GetVersionInfo(), request.GetResponseNonce(), request.GetResourceNames(), request.GetErrorDetail()}, []interface{}{response.GetVersionInfo(), response.GetNonce()})
}

func (c callbacks) OnDeltaStreamOpen(ctx context.Context, i int64, s string) error {
	//TODO implement me
	c.log.Debugf("OnDeltaStreamOpen..., i:%d, s:%s", i, s)
	return nil
}

func (c callbacks) OnDeltaStreamClosed(i int64) {
	//TODO implement me
	c.log.Debugf("OnDeltaStreamClosed..., i:%d", i)
}

func (c callbacks) OnStreamDeltaRequest(i int64, request *discoverygrpcv3.DeltaDiscoveryRequest) error {
	// c.log.Debugf("OnStreamDeltaRequest..., i:%d request:%+v", i, request)
	c.log.Debugf("OnStreamDeltaRequest..., i:%d request:%+v", i, []interface{}{request.GetResponseNonce(), request.GetResponseNonce(), request.GetResourceNamesSubscribe(), request.GetResourceNamesUnsubscribe(), request.GetErrorDetail()})
	// TODO: delta请求处理，下发ploycube配置
	return nil
}

func (c callbacks) OnStreamDeltaResponse(i int64, request *discoverygrpcv3.DeltaDiscoveryRequest, response *discoverygrpcv3.DeltaDiscoveryResponse) {
	//TODO implement me
	// c.log.Debugf("OnStreamDeltaResponse..., i:%d request:%+v, response:%+v", i, request, response)
	c.log.Debugf("OnStreamDeltaResponse..., i:%d request:%+v, response:%+v", i, []interface{}{request.GetResponseNonce(), request.GetResponseNonce(), request.GetResourceNamesSubscribe(), request.GetResourceNamesUnsubscribe(), request.GetErrorDetail()}, []interface{}{response.GetNonce(), response.GetRemovedResourceNames()})
}
