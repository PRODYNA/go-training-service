package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleWebPort(router *mux.Router) {

	router.
		Methods("GET").
		Path("/status").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			_, err := writer.Write([]byte(`{ "status": "ok"}`))
			if err != nil {
				return
			}
		})
}
