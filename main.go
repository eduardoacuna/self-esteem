package main

import (
	"github.com/eduardoacuna/self-esteem/server"
)

func main() {
	server.SetupRoutes()
	server.ListenAndServe()
}
