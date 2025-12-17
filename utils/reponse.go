package utils

import(
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, status int, data any){

	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data": data,
	})
}


func Error(w http.ResponseWriter, status int, msg string){

	w.Header().Set("Content-Type","application/json")
	
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error": msg,
	})
}