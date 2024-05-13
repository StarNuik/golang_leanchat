package server

import (
	"context"
	"net/url"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlClient struct {
	pool  *pgxpool.Pool
	close func()
}

func BuildSqlClient(url *url.URL) (*SqlClient, error) {
	conf, err := pgxpool.ParseConfig(url.String())
	if err != nil {
		return nil, err
	}
	conf.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxuuid.Register(c.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.TODO(), conf)
	if err != nil {
		return nil, err
	}
	close := func() {
		pool.Close()
	}

	err = pool.Ping(context.TODO())
	if err != nil {
		close()
		return nil, err
	}

	sc := SqlClient{
		pool:  pool,
		close: close,
	}
	return &sc, nil
}

func (sc *SqlClient) Close() {
	sc.close()
}

func (sc *SqlClient) Conn() *pgxpool.Pool {
	return sc.pool
}

// func (sc *SqlClient) GetChannelMessages(chanId *uuid.UUID, count uint) {
//
// }
