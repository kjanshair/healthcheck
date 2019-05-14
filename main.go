package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(os.Getenv("ENDPOINT"), HttpOk)

	fmt.Println("Started at :" + os.Getenv("PORT") + os.Getenv("ENDPOINT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.LoggingHandler(os.Stdout, r)))
}

func HttpOk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "OK\n")
}
