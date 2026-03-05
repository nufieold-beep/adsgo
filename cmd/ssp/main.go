
package main

import (
	"log"
	"ssp/internal/http"
)

func main() {

	app := http.NewRouter()

	log.Println("SSP production server started :8080")

	log.Fatal(app.Listen(":8080"))
}
