package main

import (
	"API_go/go_test/config"
	migration "API_go/go_test/migrations"
	"API_go/go_test/router"
)

func init() {
	config.LoadEnv()
	config.ConnectToDatabase()
	migration.Migrate()
}

func main() {
	router.Routes()
}
