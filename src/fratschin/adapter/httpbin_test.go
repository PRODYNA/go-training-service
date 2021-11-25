package adapter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestAdapterImpl_CallBackend(t *testing.T) {

	a := NewAdapter(&http.Client{
		Transport: mock{},
	})

	assert.NotNil(t, a)

	res, err := a.CallBackend()
	assert.Nil(t, err)

	fmt.Println(res)
}

type mock struct {

}

func (m mock) RoundTrip(*http.Request) (*http.Response, error) {

	b := ioutil.NopCloser(strings.NewReader(responseData))
	return &http.Response{Body: b}, nil
}


var responseData string  = `{
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
}`
