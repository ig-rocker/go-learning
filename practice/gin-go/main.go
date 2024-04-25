package main

import (
	"learngin/routes"
)

func main() {

	route := routes.Route()
	route.Run(":8080")
}
