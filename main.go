package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"

	"github.com/micro/place-srv/google"
	"github.com/micro/place-srv/handler"
	proto "github.com/micro/place-srv/proto/google"
)

func main() {
	cmd.Flags = append(cmd.Flags,
		cli.StringFlag{
			Name:   "google_api_key",
			EnvVar: "GOOGLE_API_KEY",
			Usage:  "Google maps API key",
		},
	)

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		google.Key = c.String("google_api_key")
	})

	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.place"),
	)

	proto.RegisterGoogleHandler(server.DefaultServer, &handler.Google{})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
