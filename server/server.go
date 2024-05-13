package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"starnuik/leanchat/rpc"

	"google.golang.org/grpc"
)

type ServerContext struct {
	Port        uint16
	PostgresUrl *url.URL
}

type RpcServer struct {
	rpc.UnimplementedLeanchatServer
	sql *SqlClient
}

func checkServerError(err error, where string) {
	if err != nil {
		log.Fatalf("%s error: %v\n", where, err)
	}
}

func relayError(err error, where string) error {
	log.Printf("%s error: %v\n", where, err)
	return fmt.Errorf("%s error", where)
}

func RunServer(ctx *ServerContext) {
	sql, err := BuildSqlClient(ctx.PostgresUrl)
	checkServerError(err, "sql.BuildSqlClient")
	defer sql.Close()

	tcpUrl := fmt.Sprintf(":%d", ctx.Port)
	lis, err := net.Listen("tcp", tcpUrl)
	checkServerError(err, "net.Listen")

	rs := RpcServer{
		sql: sql,
	}
	s := grpc.NewServer()
	rpc.RegisterLeanchatServer(s, &rs)

	log.Println("Server started")
	err = s.Serve(lis)
	checkServerError(err, "s.Serve")
}

func (rs *RpcServer) PeekChannel(ctx context.Context, req *rpc.PeekChannelRequest) (*rpc.PeekChannelResponse, error) {
	log.Println("PeekChannel received")
	res := rpc.PeekChannelResponse{}
	sql := rs.sql.Conn()

	row := sql.QueryRow(context.TODO(), `
		SELECT chan_name
		FROM channels
		WHERE chan_id = $1`,
		req.ChanId.Data)
	err := row.Scan(&res.ChanName)
	if err != nil {
		return nil, relayError(err, "row.Scan")
	}

	rows, err := sql.Query(context.TODO(), `
		SELECT user_name, content
		FROM messages
		WHERE chan_id = $1
		ORDER BY msg_created DESC
		LIMIT $2
	`, req.ChanId.Data, max(req.ReqCount, 255))
	if err != nil {
		return nil, relayError(err, "sql.Query")
	}
	defer rows.Close()

	msgs := []*rpc.ChatMessage{}
	for rows.Next() {
		msg := rpc.ChatMessage{}
		err = rows.Scan(&msg.UserName, &msg.MsgContent)
		if err != nil {
			return nil, relayError(err, "rows.Scan")
		}
		msgs = append(msgs, &msg)
	}
	res.Messages = msgs

	log.Println("PeekChannel responded")
	return &res, nil
}
