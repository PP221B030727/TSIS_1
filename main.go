package main

import (
	"TSIS_1/api"
	"fmt"
)

func main() {
	fmt.Println("Start server ...")
	api.SetupRouter()
	fmt.Println("Server is running on port 8080...")
}
