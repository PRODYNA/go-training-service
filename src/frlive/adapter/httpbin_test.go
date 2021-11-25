package adapter

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestIntegrationAdapter_PostData(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	a := NewAdapter(http.DefaultClient)
	data, err := a.PostData()
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestAdapter_PostData(t *testing.T) {
	a := NewAdapter(&http.Client{
		Transport: mock{},
	})
	data, err := a.PostData()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "92.73.20.61", data.Origin)
}

type mock struct{}

func (m mock) RoundTrip(*http.Request) (*http.Response, error) {

	b := ioutil.NopCloser(strings.NewReader(okRes))
	res :=  &http.Response{
		Body: b,
	}
	return res, nil
}

var okRes = `{
    "args": {},
    "data": "",
    "files": {},
    "form": {},
    "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "en-US,en;q=0.9,de;q=0.8",
        "Cache-Control": "no-cache",
        "Content-Length": "0",
        "Host": "httpbin.org",
        "Origin": "chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop",
        "Postman-Token": "30fa87fe-a7fa-b1eb-db71-3649004f3e46",
        "Sec-Ch-Ua": "\"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"",
        "Sec-Ch-Ua-Mobile": "?0",
        "Sec-Ch-Ua-Platform": "\"Linux\"",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "none",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
        "X-Amzn-Trace-Id": "Root=1-61911242-22396b201a14eee3189265a4"
    },
    "json": null,
    "origin": "92.73.20.61",
    "url": "https://httpbin.org/post"
}`
