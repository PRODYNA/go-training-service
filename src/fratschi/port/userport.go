package port

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training/fratschi/adapter"
	"net/http"
)

func HandleUserPort(router *mux.Router, adapter adapter.HttpBin) {

	router.
		Methods("GET").
		Path("/user").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			res ,err := adapter.DoBackendCall()

			if err != nil {
				writer.WriteHeader(500)
			}

			writer.WriteHeader(200)
			d, _ := json.Marshal(res)
			writer.Write(d)
		})

}
