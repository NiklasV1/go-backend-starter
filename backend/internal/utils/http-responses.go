package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeDTO[T any](w http.ResponseWriter, dto T) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		fmt.Println("Unable to parse response DTO to json")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

