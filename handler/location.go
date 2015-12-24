package handler

import (
	"github.com/micro/go-micro/errors"
	"github.com/micro/place-srv/elastic"
	proto "github.com/micro/place-srv/proto/location"

	"golang.org/x/net/context"
)

type Location struct{}

func (l *Location) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	if req.Place == nil {
		return errors.BadRequest("go.micro.sv.place.Location.Create", "Place cannot be blank")
	}

	if err := elastic.Create(req.Place); err != nil {
		return errors.InternalServerError("go.micro.srv.place.Locate.Create", err.Error())
	}

	return nil
}

func (l *Location) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("go.micro.sv.place.Location.ead", "Id cannot be blank")
	}

	pl, err := elastic.Read(req.Id)
	if err != nil && err.Error() == elastic.ErrNotFound.Error() {
		return errors.NotFound("go.micro.srv.place.Locate.Read", err.Error())
	} else if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Locate.Read", err.Error())
	}

	rsp.Place = pl

	return nil
}

func (l *Location) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	if req.Place == nil {
		return errors.BadRequest("go.micro.sv.place.Location.Update", "Place cannot be blank")
	}

	if err := elastic.Update(req.Place); err != nil {
		return errors.InternalServerError("go.micro.srv.place.Locate.Update", err.Error())
	}

	return nil
}

func (l *Location) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("go.micro.sv.place.Location.Delete", "Id cannot be blank")
	}

	err := elastic.Delete(req.Id)
	if err != nil && err.Error() == elastic.ErrNotFound.Error() {
		return nil
	} else if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Locate.Delete", err.Error())
	}

	return nil
}

func (l *Location) Search(ctx context.Context, req *proto.SearchRequest, rsp *proto.SearchResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}

	places, err := elastic.Search(req.Query, req.Limit, req.Offset)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.place.Locate.Search", err.Error())
	}
	rsp.Places = places
	return nil
}
