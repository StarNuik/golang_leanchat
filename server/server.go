package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"starnuik/leanchat/rpc"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
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

func checkRowsAffected(tag *pgconn.CommandTag, rows int64, where string) error {
	if tag.RowsAffected() != rows {
		return fmt.Errorf("%s, rows affected: %d, should be: %d", where, tag.RowsAffected(), rows)
	}
	return nil
}

func checkStringNotEmpty(str string) error {
	if str == "" {
		return fmt.Errorf("string is empty")
	}
	return nil
}

func maxLength(str string, max int) string {
	if len(str) > max {
		return str[:max]
	}
	return str
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
	res := rpc.PeekChannelResponse{}
	sql := rs.sql.pool

	row := sql.QueryRow(ctx, `
		SELECT chan_name
		FROM channels
		WHERE chan_id = $1`,
		req.ChanId.Data)
	err := row.Scan(&res.ChanName)
	if err != nil {
		return nil, relayError(err, "row.Scan")
	}

	rows, err := sql.Query(ctx, `
		SELECT user_name, content
		FROM messages
		WHERE chan_id = $1
		ORDER BY msg_created DESC
		LIMIT $2`,
		req.ChanId.Data,
		min(req.ReqCount, 255),
	)
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

	log.Println("responded to PeekChannel")
	return &res, nil
}

func messageChannel(ctx context.Context, sql *pgxpool.Pool, chanId *rpc.Uuid, name string, message string) error {
	err := checkStringNotEmpty(name)
	if err != nil {
		return relayError(err, "checkStringNotEmpty")
	}
	err = checkStringNotEmpty(message)
	if err != nil {
		return relayError(err, "checkStringNotEmpty")
	}

	tag, err := sql.Exec(ctx, `
		INSERT INTO messages
		(chan_id, user_name, content)
		VALUES
		($1, $2, $3)`,
		chanId.Data,
		maxLength(name, 63),
		message,
	)
	if err != nil {
		return relayError(err, "sql.Exec")
	}
	err = checkRowsAffected(&tag, 1, "sql.Exec")
	if err != nil {
		return relayError(err, "checkRowsAffected")
	}
	return nil
}

func (rs *RpcServer) MessageChannel(ctx context.Context, req *rpc.MessageChannelRequest) (*rpc.MessageChannelResponse, error) {

	err := messageChannel(ctx, rs.sql.pool, req.ChanId, req.Message.UserName, req.Message.MsgContent)
	if err != nil {
		return nil, err
	}

	log.Println("responded to MessageChannel")
	return &rpc.MessageChannelResponse{}, nil
}

func (rs *RpcServer) ListChannels(ctx context.Context, req *rpc.ListChannelsRequest) (*rpc.ListChannelsResponse, error) {
	sql := rs.sql.pool

	rows, err := sql.Query(ctx, `
		SELECT chan_id, chan_name
		FROM channels
		ORDER BY chan_created DESC
		LIMIT $1`,
		req.ReqCount,
	)
	if err != nil {
		return nil, relayError(err, "sql.Query")
	}
	defer rows.Close()

	chans := []*rpc.ChatChannel{}
	for rows.Next() {
		ch := rpc.ChatChannel{
			ChanId: &rpc.Uuid{},
		}
		err := rows.Scan(&ch.ChanId.Data, &ch.ChanName)
		if err != nil {
			return nil, relayError(err, "rows.Scan")
		}
		chans = append(chans, &ch)
	}

	res := rpc.ListChannelsResponse{
		Channels: chans,
	}
	log.Println("responded to ListChannels")
	return &res, nil
}

func (rs *RpcServer) CreateChannel(ctx context.Context, req *rpc.CreateChannelRequest) (*rpc.CreateChannelResponse, error) {
	res := rpc.CreateChannelResponse{
		Channel: &rpc.ChatChannel{
			ChanId: &rpc.Uuid{},
		},
	}
	sql := rs.sql.pool

	err := checkStringNotEmpty(req.ChanName)
	if err != nil {
		return nil, relayError(err, "checkStringNotEmpty")
	}

	res.Channel.ChanName = maxLength(req.ChanName, 63)
	row := sql.QueryRow(ctx, `
		INSERT INTO channels
		(chan_name)
		VALUES
		($1)
		RETURNING
		chan_id`,
		res.Channel.ChanName,
	)
	err = row.Scan(&res.Channel.ChanId.Data)
	if err != nil {
		return nil, relayError(err, "row.Scan")
	}

	err = messageChannel(ctx, sql, res.Channel.ChanId, "Leanchat", "Welcome to Leanchat!")
	if err != nil {
		return nil, err
	}

	log.Println("responded to CreateChannel")
	return &res, nil
}
