package main

import (
	"strings"

	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"

	"github.com/micro/place-srv/elastic"
	"github.com/micro/place-srv/google"
	"github.com/micro/place-srv/handler"
	proto "github.com/micro/place-srv/proto/google"
	proto2 "github.com/micro/place-srv/proto/location"
)

func main() {
	cmd.Flags = append(cmd.Flags,
		cli.StringFlag{
			Name:   "google_api_key",
			EnvVar: "GOOGLE_API_KEY",
			Usage:  "Google maps API key",
		},
		cli.StringFlag{
			Name:   "elasticsearch_hosts",
			EnvVar: "ELASTICSEARCH_HOSTS",
			Usage:  "Comma separated list of elasticsearch hosts",
			Value:  "localhost:9200",
		},
	)

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		google.Key = c.String("google_api_key")
		parts := strings.Split(c.String("elasticsearch_hosts"), ",")
		elastic.Hosts = parts
	})

	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.place"),
	)

	proto.RegisterGoogleHandler(server.DefaultServer, &handler.Google{})
	proto2.RegisterLocationHandler(server.DefaultServer, &handler.Location{})

	elastic.Init()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
