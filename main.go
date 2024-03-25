package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/{httpResponseStatusCode:[0-9]{3}}", func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)

		httpResponseStatusCode := vars["httpResponseStatusCode"]
		httpResponseStatusCodeInt, _ := strconv.ParseInt(httpResponseStatusCode, 10, 0)
		httpResponseStatusCodeDescription := http.StatusText(int(httpResponseStatusCodeInt))

		slog.Info("Received request for", "httpResponseStatusCode", httpResponseStatusCode)

		responseWriter.WriteHeader(int(httpResponseStatusCodeInt))
		responseWriter.Header().Set("Content-type", "application/text")

		responseWriter.Write([]byte(httpResponseStatusCode + ", " + httpResponseStatusCodeDescription))
	})

	router.NotFoundHandler = http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		url := request.URL.String()

		slog.Info("Received request for unknown", "url", url)

		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Header().Set("Content-type", "application/text")

		responseWriter.Write([]byte("Not found, " + url))
	})

	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "4444"
	}

	slog.Info("Listening on", "port", port)

	http.ListenAndServe(":"+port, router)

}
