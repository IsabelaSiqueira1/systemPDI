package main

import (
	"fmt"
	"net/http"

	httpgateway "github.com/IsabelaSiqueira1/systemPDI/src/internal/gateway/http"
)

// @title Queue Management API
// @version 1.0.0
// @description Backend API for managing service-counter queues (ticketing system).
// @description In-memory only (no database). Supports services, professionals, priority tickets, attendances and webhook notifications.
// @BasePath /
func main() {
	r := httpgateway.NewRouter()

	fmt.Println("listening on :8080")
	panic(http.ListenAndServe(":8080", r))
}
