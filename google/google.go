package google

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	Key    string
	Url    = "https://maps.googleapis.com/maps/api/"
	Format = "json"
)

func Do(method string, args url.Values) ([]byte, error) {
	u := Url + method + "/" + Format

	if len(Key) > 0 {
		args.Set("key", Key)
	}
	fmt.Println(args.Encode())
	rsp, err := http.Get(u + "?" + args.Encode())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
