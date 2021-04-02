package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carlosh1989/candyshop/pkg/adding"
	"github.com/carlosh1989/candyshop/pkg/http/rest"
	"github.com/carlosh1989/candyshop/pkg/reading"
	"github.com/carlosh1989/candyshop/pkg/storage"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}
	rs := reading.NewService(r)
	as := adding.NewService(r)
	fmt.Println("Starting server on port:8080")
	router := rest.InitHandlers(rs, as)
	log.Fatal(http.ListenAndServe(":8080", router))
}
