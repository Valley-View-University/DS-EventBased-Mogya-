package handlers

import (
	"encoding/json"
	"net/http"

	"../models"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set(  "Content-Type",  "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Alive{
		true,
		"Hello Kwabena",
	})

}