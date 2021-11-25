package adapter

import (
	"encoding/json"
	"github.com/prodyna/fratschin/data"
	"io/ioutil"
	"net/http"
	"strings"
)

type Adapter interface {
	CallBackend() (*data.User, error)
}

type AdapterImpl struct {
	Client *http.Client
}

func NewAdapter(client *http.Client) Adapter {
	return &AdapterImpl{
		Client: client,
	}
}

func (a AdapterImpl) CallBackend() (*data.User, error) {

	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader(`{ "ok" : true }`))
	if err != nil {
		return nil, err
	}

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &data.User{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
