package adapter

import (
	"github.com/prodyna/go-training/fratschi/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestIntegrationHttpBinImpl_DoBackendCall(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	a := NewHttpBin(config.Configuration{}, http.DefaultClient)

	res,err := a.DoBackendCall()
	assert.Nil(t, err)
	assert.NotNil(t, res.Data)

}


func TestHttpBinImpl_DoBackendCall(t *testing.T) {
	a := NewHttpBin(config.Configuration{}, &http.Client{
		Transport: Mock{},
	})

	res,err := a.DoBackendCall()
	assert.Nil(t, err)
	assert.NotNil(t, res.Data)

}


type Mock struct {}

func (m Mock) RoundTrip(*http.Request) (*http.Response, error) {
	b := ioutil.NopCloser(strings.NewReader(respone))
	res := http.Response{Body:b ,StatusCode: 200}
	return &res, nil
}


var respone = `{
  "args": {}, 
  "data": "{\n\t\"test\" : true\n}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "*/*", 
    "Cache-Control": "no-cache", 
    "Content-Length": "18", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "curl/7.74.0", 
    "X-Amzn-Trace-Id": "Root=1-6193cb71-116923410f83d84a7197a6b0"
  }, 
  "json": {
    "test": true
  }, 
  "origin": "92.73.20.61", 
  "url": "https://httpbin.org/post"
}
`
