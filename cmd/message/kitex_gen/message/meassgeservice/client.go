// Code generated by Kitex v0.4.4. DO NOT EDIT.

package meassgeservice

import (
	message "Simple-Douyin/cmd/message/kitex_gen/message"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessageAction(ctx context.Context, req *message.MessageActionRequest, callOptions ...callopt.Option) (r *message.MessageActionResponse, err error)
	MessageChat(ctx context.Context, req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMeassgeServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMeassgeServiceClient struct {
	*kClient
}

func (p *kMeassgeServiceClient) MessageAction(ctx context.Context, req *message.MessageActionRequest, callOptions ...callopt.Option) (r *message.MessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageAction(ctx, req)
}

func (p *kMeassgeServiceClient) MessageChat(ctx context.Context, req *message.MessageChatRequest, callOptions ...callopt.Option) (r *message.MessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageChat(ctx, req)
}
