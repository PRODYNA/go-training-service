package port

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/fratschin/data"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandleRest(t *testing.T) {


	rec := httptest.NewRecorder()
	router := mux.NewRouter()
	HandleRest(router, &AdapterMock{})

	router.ServeHTTP(rec, httptest.NewRequest("GET", "https://whatever/user", nil))

	assert.Equal(t, 200, rec.Code)

}

type AdapterMock struct{}


func (m AdapterMock) CallBackend() (*data.User, error) {

	return &data.User{
		URL: "testurl",
	}, nil
}


