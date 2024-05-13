package client

import (
	"fmt"
	"net/url"
	"starnuik/leanchat/rpc"
	"starnuik/leanchat/server"
)

func (cmd *runServerCmd) Run() error {
	urlStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cmd.PostgresUser,
		cmd.PostgresPassword,
		cmd.PostgresHost,
		cmd.PostgresPort,
		cmd.PostgresDb,
	)
	url, err := url.Parse(urlStr)
	checkClientError(err, "url.Parse")

	ctx := server.ServerContext{
		Port:        cmd.Port,
		PostgresUrl: url,
	}
	server.RunServer(&ctx)
	return nil
}

func (cmd *peekChannelCmd) Run() error {
	rc, err := BuildRpcClient(cmd.ServerUrl)
	checkClientError(err, "rpc.BuildRpcClient")
	defer rc.Close()

	req := rpc.PeekChannelRequest{
		ChanId:   rpc.PackUuid(cmd.ChannelId),
		ReqCount: uint32(cmd.Count),
	}
	res, err := rc.Conn().PeekChannel(rc.Context, &req)
	checkClientError(err, "rc.PeekChannel")

	msgs := res.Messages
	fmt.Printf("peek channel[ %s ]\n", res.ChanName)
	for idx := len(msgs) - 1; idx >= 0; idx-- {
		msg := msgs[idx]
		fmt.Printf("  user[ %s ]: %s\n", msg.UserName, msg.MsgContent)
	}

	return nil
}
