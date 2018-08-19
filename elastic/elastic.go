package elastic

import (
	"encoding/json"
	"errors"
	"fmt"
	lib "github.com/mattbaird/elastigo/lib"
	proto "github.com/microhq/place-srv/proto/location"
)

var (
	ErrNotFound = errors.New("record not found")

	Index = "places"
	Type  = "place"
	Hosts []string
	conn  *lib.Conn
)

func Init() {
	conn = lib.NewConn()
	conn.SetHosts(Hosts)
}

func Create(pl *proto.Place) error {
	_, err := conn.Index(Index, Type, pl.Id, nil, pl)
	return err
}

func Update(pl *proto.Place) error {
	_, err := conn.Index(Index, Type, pl.Id, nil, pl)
	return err
}

func Read(id string) (*proto.Place, error) {
	r, err := conn.Get(Index, Type, id, nil)
	if err != nil {
		return nil, err
	}

	var pl *proto.Place

	if err := json.Unmarshal(*r.Source, &pl); err != nil {
		return nil, err
	}

	return pl, nil
}

func Delete(id string) error {
	_, err := conn.Delete(Index, Type, id, nil)
	return err
}

func Search(query string, limit, offset uint64) ([]*proto.Place, error) {
	if len(query) == 0 {
		query = "*"
	}
	size := fmt.Sprintf("%d", limit)
	from := fmt.Sprintf("%d", offset)

	out, err := lib.Search(Index).Type(Type).Size(size).From(from).Search(query).Result(conn)
	if err != nil {
		return nil, err
	}

	var places []*proto.Place
	for _, hit := range out.Hits.Hits {
		var pl *proto.Place
		if err := json.Unmarshal(*hit.Source, &pl); err != nil {
			return nil, err
		}
		places = append(places, pl)
	}

	return places, nil
}
