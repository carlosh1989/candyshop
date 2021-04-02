package rest

import (
	"encoding/json"
	"net/http"

	"github.com/carlosh1989/candyshop/pkg/adding"
)

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		id, err := as.AddCandy(nc)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		nc.Id = id
		encoder := json.NewEncoder(w)
		encoder.Encode(nc)
	}
}
