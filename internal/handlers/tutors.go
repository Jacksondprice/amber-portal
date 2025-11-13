package handlers

import (
	"encoding/json"
	"net/http"
)

type Tutor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListTutors(w http.ResponseWriter, _ *http.Request) {
	items := []Tutor{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
