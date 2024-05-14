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

type createChannelCmd struct{}
type runServerCmd struct {
	Port             uint16 `arg:""`
	PostgresUser     string `arg:"" env:"POSTGRES_USER"`
	PostgresPassword string `arg:"" env:"POSTGRES_PASSWORD"`
	PostgresHost     string `arg:"" env:"POSTGRES_HOST"`
	PostgresPort     uint16 `arg:"" env:"POSTGRES_PORT"`
	PostgresDb       string `arg:"" env:"POSTGRES_DB"`
}
