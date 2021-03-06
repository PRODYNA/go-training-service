package adapter

import (
	"encoding/json"
	"errors"
	"github.com/prodyna/go-training/fratschi/config"
	"github.com/prodyna/go-training/fratschi/data"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpBin interface {
	DoBackendCall() (*data.Result, error)
}

type httpBinImpl struct {
	Config config.Configuration
	Client *http.Client
}

func NewHttpBin(cfg config.Configuration, client *http.Client) HttpBin {
	return &httpBinImpl{
		Config: cfg,
		Client: client,
	}
}

func (h httpBinImpl) DoBackendCall() (*data.Result, error) {

	req, err := http.NewRequest("POST","https://httpbin.org/post", strings.NewReader(`{ "ok" : true }`))
	if err != nil {
		return nil, err
	}

	res, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("got a 500")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := data.Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
