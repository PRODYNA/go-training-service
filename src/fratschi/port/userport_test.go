package port

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training/fratschi/adapter"
	"github.com/prodyna/go-training/fratschi/config"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleUserPort(t *testing.T) {

	rec := httptest.NewRecorder()
	router := mux.NewRouter()
	HandleUserPort(router, adapter.NewHttpBin(config.Configuration{}, &http.Client{
		Transport: Mock{},
	} ))

	router.ServeHTTP(rec, httptest.NewRequest("GET", "https://any:1234/user", strings.NewReader(`{ }`)))

}



type Mock struct {}

func (m Mock) RoundTrip(*http.Request) (*http.Response, error) {
	b := ioutil.NopCloser(strings.NewReader(respone))
	res := http.Response{Body:b }
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
