package port

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prodyna/fratschin/adapter"
	"net/http"
)

func HandleRest(router *mux.Router, a adapter.Adapter) {

	router.
		Methods("GET").
		Path("/user").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			user, err := a.CallBackend()
			if err != nil {
				writer.WriteHeader(500)
				return
			}

			bytes, err := json.Marshal(user)
			if err != nil {
				writer.WriteHeader(500)
				return
			}

			writer.Write(bytes)

			writer.WriteHeader(200)
		})
}
