package client

import (
	"context"
	"fmt"
	"net/url"
	"github.com/starnuik/golang_leanchat/rpc"
	"github.com/starnuik/golang_leanchat/server"

	"github.com/gofrs/uuid/v5"
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

func buildRpcClient(url *url.URL) *RpcClient {
	rc, err := BuildRpcClient(url)
	checkClientError(err, "rpc.BuildRpcClient")
	return rc
}

func peekChannel(rc *RpcClient, channelId *uuid.UUID, count uint8) {
	req := rpc.PeekChannelRequest{
		ChanId:   rpc.PackUuid(channelId),
		ReqCount: uint32(count),
	}
	res, err := rc.rpc.PeekChannel(rc.Context, &req)
	checkClientError(err, "rc.PeekChannel")

	msgs := res.Messages
	fmt.Printf("peek channel[ %s ]\n", res.ChanName)
	for idx := len(msgs) - 1; idx >= 0; idx-- {
		msg := msgs[idx]
		fmt.Printf("  user[ %s ]: %s\n", msg.UserName, msg.MsgContent)
	}
}

func (cmd *peekChannelCmd) Run() error {
	rc := buildRpcClient(cmd.ServerUrl)
	defer rc.Close()

	peekChannel(rc, cmd.ChannelId, cmd.Count)

	return nil
}

func (cmd *messageChannelCmd) Run() error {
	rc := buildRpcClient(cmd.ServerUrl)
	defer rc.Close()

	req := rpc.MessageChannelRequest{
		ChanId: rpc.PackUuid(cmd.ChannelId),
		Message: &rpc.ChatMessage{
			UserName:   cmd.UserName,
			MsgContent: cmd.Message,
		},
	}
	_, err := rc.rpc.MessageChannel(rc.Context, &req)
	checkClientError(err, "rc.MessageChannel")

	peekChannel(rc, cmd.ChannelId, cmd.Count)

	return nil
}

func (cmd *listChannelsCmd) Run() error {
	rc := buildRpcClient(cmd.ServerUrl)
	defer rc.Close()

	req := rpc.ListChannelsRequest{
		ReqCount: uint32(cmd.Count),
	}
	res, err := rc.rpc.ListChannels(context.TODO(), &req)
	checkClientError(err, "rc.SearchChannels")

	fmt.Printf("search channels\n")
	chs := res.Channels
	for idx := len(chs) - 1; idx >= 0; idx-- {
		ch := chs[idx]
		fmt.Printf("  channel[ %s ], channel_id[ %s ]\n", ch.ChanName, ch.ChanId.Data)
	}
	return nil
}

func (cmd *createChannelCmd) Run() error {
	rc := buildRpcClient(cmd.ServerUrl)
	defer rc.Close()

	req := rpc.CreateChannelRequest{
		ChanName: cmd.ChannelName,
	}
	res, err := rc.rpc.CreateChannel(context.TODO(), &req)
	checkClientError(err, "rc.CreateChannel")

	fmt.Printf("created channel\nchannel[ %s ], channel_id[ %s ]\n", res.Channel.ChanName, res.Channel.ChanId.Data)
	return nil
}
