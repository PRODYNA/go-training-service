package service

import (
	"github.com/prodyna/frlive/adapter"
	"github.com/prodyna/frlive/data"
	_ "github.com/prodyna/frlive/data"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ServiceImpl struct {
	Adapter adapter.Adapter
}

type Service interface {
	CreateUser(userid string) *data.User
	Call() (*data.Result, error)
}

func NewService(client *http.Client, adapter adapter.Adapter) *ServiceImpl {
	log.Info().Msg("creating new service")

	return &ServiceImpl{
		Adapter: adapter,
	}
}

func (s ServiceImpl) CreateUser(userid string) *data.User {

	return &data.User{
		Name: "Frank",
		Age:  49,
	}
}

func (s ServiceImpl) Call() (*data.Result, error) {
	return s.Adapter.PostData()
}
