package client

import (
	"net/url"

	"github.com/gofrs/uuid/v5"
)

type Cli struct {
	Peek    peekChannelCmd    `cmd:""`
	Message messageChannelCmd `cmd:""`
	Create  createChannelCmd  `cmd:""`
	List    listChannelsCmd   `cmd:""`
	Serve   runServerCmd      `cmd:""`
}

type peekChannelCmd struct {
	ServerUrl *url.URL   `arg:""`
	ChannelId *uuid.UUID `arg:""`
	Count     uint8      `default:"31"`
}
type messageChannelCmd struct {
	ServerUrl *url.URL   `arg:""`
	ChannelId *uuid.UUID `arg:""`
	UserName  string     `arg:""`
	Message   string     `arg:""`
	Count     uint8      `default:"31"`
}
type listChannelsCmd struct {
	ServerUrl *url.URL `arg:""`
	Count     uint8    `default:"31"`
}
type createChannelCmd struct {
	ServerUrl   *url.URL `arg:""`
	ChannelName string   `arg:""`
}
type runServerCmd struct {
	Port             uint16 `arg:""`
	PostgresUser     string `arg:"" env:"PGUSER"`
	PostgresPassword string `arg:"" env:"PGPASSWORD"`
	PostgresHost     string `arg:"" env:"PGHOST"`
	PostgresPort     uint16 `arg:"" env:"PGPORT"`
	PostgresDb       string `arg:"" env:"PGDATABASE"`
}
