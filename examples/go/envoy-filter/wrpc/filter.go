package main

import (
	"context"
	"net/http"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"wrpc.io/examples/go/envoy-filter/bindings/wrpc_examples/hello/handler"
	wrpc "wrpc.io/go"
)

type filter struct {
	api.PassThroughStreamFilter

	callbacks api.FilterCallbackHandler
	client    wrpc.Invoker
}

func NewFilter(client wrpc.Invoker) *filter {
	return &filter{client: client}
}

func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	path, _ := header.Get(":path")
	method, _ := header.Get(":method")

	if method == http.MethodGet && path == "/hello" {
		return f.handleHelloRequest()
	}

	return api.Continue
}

func (f *filter) handleHelloRequest() api.StatusType {
	greeting, err := handler.Hello(context.Background(), f.client)
	if err != nil {
		api.LogErrorf("failed to call `wrpc-examples:hello/handler.hello`: %v", err)
		return f.sendLocalReply(http.StatusInternalServerError, "Internal Server Error", nil)
	}

	responseHeaders := map[string][]string{
		"Content-Type": {"text/plain"},
	}
	return f.sendLocalReply(http.StatusOK, greeting, responseHeaders)
}

func (f *filter) sendLocalReply(responseCode int, bodyText string, headers map[string][]string) api.StatusType {
	f.callbacks.DecoderFilterCallbacks().SendLocalReply(responseCode, bodyText, headers, 0, "")
	// Remember to return LocalReply when the request is replied locally
	return api.LocalReply
}
