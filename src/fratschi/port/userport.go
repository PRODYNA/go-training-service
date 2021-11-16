package port

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training/fratschi/adapter"
	"net/http"
)

func HandleUserPort(router *mux.Router, adapter adapter.HttpBin) {

	router.
		Methods("GET").
		Path("/user").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			adapter.DoBackendCall()

			writer.WriteHeader(200)
			writer.Write([]byte(`{ "ok" : true }`))
		})

}
