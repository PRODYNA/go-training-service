package adapter

import (
	"encoding/json"
	"github.com/prodyna/go-training/geoo/config"
	"github.com/prodyna/go-training/geoo/data"
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
	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader(`{"ok": true}`))
	if err != nil {
		return nil, err
	}

	res, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := data.Result{}
	json.Unmarshal(body, &result)

	return &result, nil
}

