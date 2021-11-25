package port

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prodyna/frlive/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserPort struct {
	Service *service.ServiceImpl
}

func NewUserPort(s *service.ServiceImpl) *UserPort {
	log.Info().Msg("starting service")
	return &UserPort{
		s,
	}
}

func (a UserPort) HandleHttp(router *mux.Router) {

	router.
		Methods("GET").
		Path("/user/{userid}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			userid := vars["userid"]

			user := a.Service.CreateUser(userid)

			data, err := json.Marshal(user)
			if err != nil {
				log.Err(err).Msg("Cant create user")
				w.WriteHeader(500)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(data)
		})

	router.
		Methods("POST").
		Path("/user").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res, err := a.Service.Call()

			data, err := json.Marshal(res)
			if err != nil {
				log.Err(err).Msg("Cant call service")
				w.WriteHeader(500)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(data)
		})

}
