package main

import (
	"stripe_api/server"
)

func main() {
	sv := server.Register()

	sv.Start()
}
