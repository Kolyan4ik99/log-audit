package client

import (
	"context"

	"github.com/Kolyan4ik99/log-audit/pkg/rpc"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	client rpc.NotificationClient
}

func NewNotificationClient(client *grpc.ClientConn) *NotificationClient {
	return &NotificationClient{client: rpc.NewNotificationClient(client)}
}

func (n *NotificationClient) Send(ctx context.Context, r *rpc.Request) (*rpc.Response, error) {
	return n.client.Send(ctx, r)
}
