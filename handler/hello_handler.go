package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadtheswe/carbon-footprint-api/entity"
)

func HelloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := entity.Message{
		Status:  "success",
		Data:    "Hello, World!",
		Message: "Request processed successfully",
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}
