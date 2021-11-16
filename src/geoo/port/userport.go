package port

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training/geoo/adapter"
	"net/http"
)

func HandleUserPort(router *mux.Router, adapter adapter.HttpBin) {
	router.
		Methods("GET").
		Path("/user").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			res, err := adapter.DoBackendCall()
			if err != nil {
				return
			}

			json, err := json.Marshal(&res)
			if err != nil {
				return
			}

			writer.WriteHeader(200)
			writer.Write(json)
		})
}
