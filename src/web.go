package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Body string
}

func main() {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&Message{
				Body: "Hello World!",
			})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), &handler))
}
