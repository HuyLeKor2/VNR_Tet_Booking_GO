package main

import "github.com/huyle/go-ecommerce-backend-api/internal/initialize"

func main() {

	// routers := routers.NewRouter()

	// // Start the server on port 8080
	// if err := routers.Run(":8080"); err != nil {
	// 	log.Fatalf("Failed to run server: %v", err)
	// }
	initialize.Run()
}
