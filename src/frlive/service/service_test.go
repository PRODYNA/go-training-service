package service

import (

	"github.com/prodyna/frlive/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestService_Call(t *testing.T) {

	s := NewService(http.DefaultClient, &Mock{})

	res, err := s.Call()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "https://testurl.test", res.URL)


}


type Mock struct {}

func (m Mock) PostData() (*data.Result, error) {

	return &data.Result{
		URL: "https://testurl.test",
	},nil
}
