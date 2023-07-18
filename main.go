package main

import (
	"flag"

	"nats_example/baselib/app"
	"nats_example/eventstores/server"
)

func main() {
	flag.Parse()
	instance := server.NewEventstoreServer()
	app.DoMainAppInstance(instance)
}
