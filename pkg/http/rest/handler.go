package rest

import (
	"github.com/carlosh1989/candyshop/pkg/adding"
	"github.com/carlosh1989/candyshop/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(r reading.Service, as adding.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandieshandler(r)).Methods("GET")

	//Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")

	return router
}
