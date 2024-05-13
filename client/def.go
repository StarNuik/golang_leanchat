package client

import (
	"net/url"

	"github.com/gofrs/uuid/v5"
)

type Cli struct {
	Create  createChannelCmd  `cmd:""`
	Search  searchChannelsCmd `cmd:""`
	Peek    peekChannelCmd    `cmd:""`
	Message messageChannelCmd `cmd:""`
	Serve   runServerCmd      `cmd:""`
}

type peekChannelCmd struct {
	ServerUrl *url.URL   `arg:""`
	ChannelId *uuid.UUID `arg:""`
	Count     uint8      `default:"31"`
}
type createChannelCmd struct{}
type searchChannelsCmd struct{}
type messageChannelCmd struct{}
type runServerCmd struct {
	Port             uint16 `arg:""`
	PostgresUser     string `arg:"" env:"POSTGRES_USER"`
	PostgresPassword string `arg:"" env:"POSTGRES_PASSWORD"`
	PostgresHost     string `arg:"" env:"POSTGRES_HOST"`
	PostgresPort     uint16 `arg:"" env:"POSTGRES_PORT"`
	PostgresDb       string `arg:"" env:"POSTGRES_DB"`
}
