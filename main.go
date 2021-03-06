package main

import (
	"strings"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"log"

	"github.com/microhq/place-srv/elastic"
	"github.com/microhq/place-srv/google"
	"github.com/microhq/place-srv/handler"
	proto "github.com/microhq/place-srv/proto/google"
	proto2 "github.com/microhq/place-srv/proto/location"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.place"),
		micro.Flags(
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
		),
		micro.Action(func(c *cli.Context) {
			google.Key = c.String("google_api_key")
			parts := strings.Split(c.String("elasticsearch_hosts"), ",")
			elastic.Hosts = parts
		}),
	)

	service.Init()
	elastic.Init()

	proto.RegisterGoogleHandler(service.Server(), &handler.Google{})
	proto2.RegisterLocationHandler(service.Server(), &handler.Location{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
