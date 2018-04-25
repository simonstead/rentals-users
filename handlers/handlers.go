package handlers

import (
	"encoding/json"
	"github.com/simonstead/rentals-users/structs"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &structs.JsonRepsonse{Msg: "success", Data: "this is some data for you"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
