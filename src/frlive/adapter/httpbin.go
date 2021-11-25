package adapter

import (
	"encoding/json"
	"github.com/prodyna/frlive/data"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
)

type AdapterImpl struct {
	Client *http.Client
}

type Adapter interface {
	PostData() (*data.Result, error)
}

func NewAdapter(client *http.Client) Adapter {
	return &AdapterImpl{
		Client: client,
	}
}

func (a AdapterImpl) PostData() (*data.Result, error) {

	reqData := `{ "test" : "ok" }`
	reqBody := ioutil.NopCloser(strings.NewReader(reqData))
	req, err := http.NewRequest("POST", "https://httpbin.org/post", reqBody)

	if err != nil {
		log.Err(err).Msg("cant create http request")
		return nil, err
	}

	res, err := a.Client.Do(req)

	if err != nil {
		log.Err(err).Msg("cant call httbin")
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Err(err).Msg("cant read response body")
		return nil, err
	}

	data := &data.Result{}
	err = json.Unmarshal(resBody, data)
	if err != nil {
		log.Err(err).Msg("cant unmarshal response body")
		return nil, err
	}

	return data, nil

}
