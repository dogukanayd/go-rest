package main

import (
	"github.com/dogukanayd/go-rest/app/api"
	"log"
)

func main() {
	err := api.Start().Run(":" + "8080")

	if err != nil {
		log.Println(err)
	}
}
