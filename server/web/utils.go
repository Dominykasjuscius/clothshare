package web

import (
	"encoding/json"
	"net/http"
)

// JSON will respond to the client as "application/json". Any provided data will
// be marshaled as JSON.
func JSON(w http.ResponseWriter, status int, data interface{}) {
	output, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if data != nil {
		w.Write(output)
	}
}

func IMAGE(w http.ResponseWriter, status int, image []byte) {
	w.Header().Set("Content-Type", "image/jpeg")

	w.WriteHeader(status)

	if image != nil {
		w.Write(image)
	}
}
