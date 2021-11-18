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
			_, err := writer.Write([]byte(`{ "status": "up"}`))
			if err != nil {
				return
			}
		})

	router.
		Methods("GET").
		Path("/").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			_, err := writer.Write([]byte(html))
			if err != nil {
				return
			}
		})
}

var html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Go - Hello World</title>
		<style>
			body {
				background-color: #545454;
			}
			h1 {
				display: flex;
				justify-content: center;
				align-items: center;
				height: 90vh;
				color: white;
			}
		</style>
	</head>
	<body>
		<h1>
		Hello World, written in Go.
		</h1>
	</body>
</html>
`
