package handler

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/micro/place-srv/google"
	proto "github.com/micro/place-srv/proto/google"

	"golang.org/x/net/context"
)

type Google struct{}

func (g *Google) NearBySearch(ctx context.Context, req *proto.NearBySearchRequest, rsp *proto.NearBySearchResponse) error {
	u := url.Values{}

	if req.Location != nil {
		u.Set("location", fmt.Sprintf("%.6f,%.6f", req.Location.Lat, req.Location.Lng))
	}
	if req.Radius > 0 {
		u.Set("radius", fmt.Sprintf("%d", req.Radius))
	}
	if len(req.RankBy) > 0 {
		u.Set("rankby", req.RankBy)
	}
	if len(req.Keyword) > 0 {
		u.Set("keyword", req.Keyword)
	}
	if req.Price != nil {
		u.Set("minprice", fmt.Sprintf("%d", req.Price.MinPrice))
		u.Set("maxprice", fmt.Sprintf("%d", req.Price.MaxPrice))
	}
	if len(req.Language) > 0 {
		u.Set("language", req.Language)
	}
	if len(req.Name) > 0 {
		u.Set("name", req.Name)
	}
	if req.OpenNow {
		u.Set("opennow", fmt.Sprintf("%t", req.OpenNow))
	}
	if len(req.Types) > 0 {
		u.Set("types", strings.Join(req.Types, "|"))
	}
	if len(req.PageToken) > 0 {
		u.Set("pagetoken", req.PageToken)
	}
	if req.ZagatSelected {
		u.Set("zagatselected", "")
	}
	b, err := google.Do("place/nearbysearch", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Google.Geocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}

func (g *Google) TextSearch(ctx context.Context, req *proto.TextSearchRequest, rsp *proto.TextSearchResponse) error {
	u := url.Values{}

	if req.Location != nil {
		u.Set("location", fmt.Sprintf("%.6f,%.6f", req.Location.Lat, req.Location.Lng))
	}
	if req.Radius > 0 {
		u.Set("radius", fmt.Sprintf("%d", req.Radius))
	}
	if len(req.Query) > 0 {
		u.Set("query", req.Query)
	}
	if req.Price != nil {
		u.Set("minprice", fmt.Sprintf("%d", req.Price.MinPrice))
		u.Set("maxprice", fmt.Sprintf("%d", req.Price.MaxPrice))
	}
	if len(req.Language) > 0 {
		u.Set("language", req.Language)
	}
	if req.OpenNow {
		u.Set("opennow", fmt.Sprintf("%t", req.OpenNow))
	}
	if len(req.Types) > 0 {
		u.Set("types", strings.Join(req.Types, "|"))
	}
	if len(req.PageToken) > 0 {
		u.Set("pagetoken", req.PageToken)
	}
	if req.ZagatSelected {
		u.Set("zagatselected", "")
	}

	b, err := google.Do("place/textsearch", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Google.ReverseGeocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}

func (g *Google) RadarSearch(ctx context.Context, req *proto.RadarSearchRequest, rsp *proto.RadarSearchResponse) error {
	u := url.Values{}

	if req.Location != nil {
		u.Set("location", fmt.Sprintf("%.6f,%.6f", req.Location.Lat, req.Location.Lng))
	}
	if req.Radius > 0 {
		u.Set("radius", fmt.Sprintf("%d", req.Radius))
	}
	if len(req.Keyword) > 0 {
		u.Set("keyword", req.Keyword)
	}
	if req.Price != nil {
		u.Set("minprice", fmt.Sprintf("%d", req.Price.MinPrice))
		u.Set("maxprice", fmt.Sprintf("%d", req.Price.MaxPrice))
	}
	if len(req.Name) > 0 {
		u.Set("name", req.Name)
	}
	if req.OpenNow {
		u.Set("opennow", fmt.Sprintf("%t", req.OpenNow))
	}
	if len(req.Types) > 0 {
		u.Set("types", strings.Join(req.Types, "|"))
	}
	if req.ZagatSelected {
		u.Set("zagatselected", "")
	}

	b, err := google.Do("place/radarsearch", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Google.ReverseGeocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
