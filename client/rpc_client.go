package client

import (
	"context"
	"net/url"
	"starnuik/leanchat/rpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RpcClient struct {
	Context context.Context

	rpc   rpc.LeanchatClient
	close func()
}

const timeout = 5 * time.Second

func BuildRpcClient(url *url.URL) (*RpcClient, error) {
	cred := insecure.NewCredentials()
	conn, err := grpc.Dial(url.String(), grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	rpc := rpc.NewLeanchatClient(conn)

	close := func() {
		cancel()
		conn.Close()
	}

	rc := &RpcClient{
		Context: ctx,
		rpc:     rpc,
		close:   close,
	}
	return rc, nil
}

func (rc *RpcClient) Close() {
	rc.close()
}
