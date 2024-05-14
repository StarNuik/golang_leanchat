package main

import (
	"log"
	"os"

	"github.com/starnuik/golang_leanchat/client"

	"github.com/alecthomas/kong"
)

func main() {
	log.SetOutput(os.Stdout)
	// rpc := client.BuildRpcClient(&session.ServerUrl)

	cli := client.Cli{}
	ctx := kong.Parse(&cli,
		kong.Name("leanchat"),
		kong.UsageOnError(),
	)
	ctx.Run( /* rpc */ )
}
