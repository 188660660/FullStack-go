// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/App.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/App.proto
	v1/BaseInfo.proto
	v1/Feed.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// =============
// App Interface
// =============

type App interface {
	// * 关注接口
	// 【粉版APP5.31 关注排序改版】App首页我的关注部分【全量数据】
	LiveHomePage(context.Context, *AppLiveHomePageReq) (*AppLiveHomePageResp, error)
}

// ===================
// App Live Rpc Client
// ===================

type appRpcClient struct {
	client *liverpc.Client
}

// NewAppRpcClient creates a Rpc client that implements the App interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewAppRpcClient(client *liverpc.Client) App {
	return &appRpcClient{
		client: client,
	}
}

func (c *appRpcClient) LiveHomePage(ctx context.Context, in *AppLiveHomePageReq) (*AppLiveHomePageResp, error) {
	out := new(AppLiveHomePageResp)
	err := doRpcRequest(ctx, c.client, 1, "App.LiveHomePage", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRpcRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message) (err error) {
	err = client.Call(ctx, version, method, in, out)
	return
}
