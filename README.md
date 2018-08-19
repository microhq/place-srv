# Place Server

Place server is a way of storing and retrieving places. It's backed by elasticsearch.

The place server also implements the Google Maps Place search API.

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Get your an API key from Google - [doc](https://developers.google.com/places/web-service/get-api-key)

4. Download and start the service

	```shell
	go get github.com/microhq/place-srv
	place-srv --google_api_key=YOUR_API_TOKEN --elasticsearch_hosts=localhost:9200
	```

	OR as a docker container

	```shell
	docker run microhq/place-srv --google_api_key=YOUR_API_TOKEN --elasticsearch_hosts=ELASTICSEARCH_HOST_ADDRESS --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The Place API

### Create Place
```shell
micro query go.micro.srv.place Location.Create '{"place": {"id": "38826c74-7790-4107-ad79-fe317f47760c", "name":"Sticks N Sushi", "iata": "LON", "address": "11 Henrietta St, London WC2E 8PY, United Kingdom", "tags": ["sushi"]}}'
{}
```

### Update Place
```shell
micro query go.micro.srv.place Location.Create '{"place": {"id": "38826c74-7790-4107-ad79-fe317f47760c", "name":"Sticks N Sushi", "iata": "LON", "address": "11 Henrietta St, London WC2E 8PY, United Kingdom", "description": "Japanese restaurant based on a renowned Danish concept, offering sushi and grilled yakitori skewers.", "phone": "+44 20 3141 8810", "location": {"lat": 51.5086542, "lng": -0.1060743}, "tags": ["sushi"]}}'
```


### Read Place
```shell
micro query go.micro.srv.place Location.Read '{"id": "38826c74-7790-4107-ad79-fe317f47760c"}'
{
	"place": {
		"address": "11 Henrietta St, London WC2E 8PY, United Kingdom",
		"description": "Japanese restaurant based on a renowned Danish concept, offering sushi and grilled yakitori skewers.",
		"iata": "LON",
		"id": "38826c74-7790-4107-ad79-fe317f47760c",
		"location": {
			"lat": 51.5086542,
			"lng": -0.1060743
		},
		"name": "Sticks N Sushi",
		"phone": "+44 20 3141 8810",
		"tags": [
			"sushi"
		]
	}
}
```

### Delete Place
```shell
micro query go.micro.srv.place Location.Delete '{"id": "38826c74-7790-4107-ad79-fe317f47760c"}'
{}
```

### Search Places
```shell
micro query go.micro.srv.place Location.Search '{"query": "tags: sushi AND iata: LON"}'
{
	"places": [
		{
			"address": "11 Henrietta St, London WC2E 8PY, United Kingdom",
			"description": "Japanese restaurant based on a renowned Danish concept, offering sushi and grilled yakitori skewers.",
			"iata": "LON",
			"id": "38826c74-7790-4107-ad79-fe317f47760c",
			"location": {
				"lat": 51.5086542,
				"lng": -0.1060743
			},
			"name": "Sticks N Sushi",
			"phone": "+44 20 3141 8810",
			"tags": [
				"sushi"
			]
		}
	]
}
```

## Google Place Search API

Place server implements the [Google Place Search API](https://developers.google.com/places/web-service/search) as RPC.

### NearBySearch
```shell
micro query go.micro.srv.place Google.NearBySearch '{"location": {"lat": -33.8670522, "lng": 151.1957362}, "radius": 500, "types": ["food"], "name": "cruise"}'
{
	"results": [
		{
			"geometry": {
				"location": {
					"lat": -33.8687895,
					"lng": 151.1942171
				}
			},
			"icon": "https://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
			"id": "21a0b251c9b8392186142c798263e289fe45b4aa",
			"name": "Rhythmboat Cruises",
			"photos": [
				{
					"height": 480,
					"html_attributions": [
						"\u003ca href=\"https://maps.google.com/maps/contrib/104066891898402903288/photos\"\u003eRhythmboat Cruises\u003c/a\u003e"
					],
					"photo_reference": "CmRdAAAAoVZ9PZykKWHs9D2kKx_wAX6_QlWTKuPxm_gIe9PCqAxgS9A6o7qJ7kQyz0OcowLTbx84inm4rIPjHZJ1vN1jorEeahOgj9-lKOhHPTdXIQYbydJOe1LAWsn78nEqnkJ7EhBGhjNIHMpilDhUwbIxj1OkGhQyUzAyfFg7crymY-WzOkD7LxpIkg",
					"width": 640
				}
			],
			"place_id": "ChIJyWEHuEmuEmsRm9hTkapTCrk",
			"reference": "CnRmAAAA2NxJ5Y6SNMxaSY3bYUuhb10ibA6nIRRBr__mjQ62WnhYbUspSjxsJ7ua7LyG2548IxODWBHrdB5gS-HqImtYjtTfgvzyk8GXVCvJhhJmYwaZiECEdkIiWw7PZIlga4kpLk6xnSOOgkZnEx0geCt_DRIQxQfmSSc0-WafPtiXaE_trRoUPlt0VZmDvJp5mC8LMsrsUDcbD8E",
			"types": [
				"restaurant",
				"food",
				"point_of_interest",
				"establishment"
			],
			"vicinity": "Pyrmont Bay Wharf (Near Australia Maritime Museum), Pyrmont, NSW 2009"
		},
...
}
```

### TextSearch
```shell
micro query go.micro.srv.place Google.TextSearch '{"query": "restaurants in sydney"}'
{
	"results": [
		{
			"formatted_address": "529 Kent St, Sydney NSW 2000, Australia",
			"geometry": {
				"location": {
					"lat": -33.875154,
					"lng": 151.204976
				}
			},
			"icon": "https://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
			"id": "827f1ac561d72ec25897df088199315f7cbbc8ed",
			"name": "Tetsuya's Restaurant",
			"opening_hours": {},
			"photos": [
				{
					"height": 2852,
					"html_attributions": [
						"\u003ca href=\"https://maps.google.com/maps/contrib/105395521304326551280/photos\"\u003eTetsuya\u0026#39;s Restaurant\u003c/a\u003e"
					],
					"photo_reference": "CmRdAAAAQuizaUFKB31Q2VRAn82XDfLtXzzH5_2wnK5uTdZSzllCoQODwrOyavkG3_7bmWBJqK13xgLjM2VcMIr0H7AH8P9LCFH9nKw1rrApnmQJpwfb2mn5u0A0Ew3JKWEOSKipEhAawaJz9HLFau4bNZyWtrQNGhQ7bMs8c8a8e3LUwA1P2utK0y2k0g",
					"width": 4296
				}
			],
			"place_id": "ChIJxXSgfDyuEmsR3X5VXGjBkFg",
			"reference": "CnRnAAAA_Gj5CU_0aBucd0R4C2TCzi0AliRSCc-Amq6CfM3MOr-cKJrMXmrENlXQszTKSsnk1oUsrKazvLoDRwqyw8TywJ8Csq8NV4KXVaYjbgEnkgRq1TfC2yXxkIutqJnRGQPJW5F1X0BYGGO2Tw6ML3x86xIQ2g_YZyRmhLnIeKw_-ZIu7BoUOceReaCJ0pE47aXJfeYormQR6sg",
			"types": [
				"restaurant",
				"food",
				"point_of_interest",
				"establishment"
			]
		},
...
}
```

### RadarSearch
```shell
micro query go.micro.srv.place Google.RadarSearch '{"location": {"lat": -33.8670522, "lng": 151.1957362}, "radius": 500, "types": ["food"], "name": "cruise"}'
{
	"results": [
		{
			"geometry": {
				"location": {
					"lat": -33.8687895,
					"lng": 151.1942171
				}
			},
			"place_id": "ChIJyWEHuEmuEmsRm9hTkapTCrk"
		},
		{
			"geometry": {
				"location": {
					"lat": -33.867591,
					"lng": 151.201196
				}
			},
			"place_id": "ChIJrTLr-GyuEmsRBfy61i59si0"
		},
		{
			"geometry": {
				"location": {
					"lat": -33.8709434,
					"lng": 151.1903114
				}
			},
			"place_id": "ChIJLfySpTOuEmsRPCRKrzl8ZEY"
		}
	],
	"status": "OK"
}
```

