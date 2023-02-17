package main

import (
	"API_go/go_test/config"
	"API_go/go_test/datafixtures"
	migration "API_go/go_test/migrations"
	"API_go/go_test/router"
)

func init() {
	config.LoadEnv()
	config.ConnectToDatabase()
	migration.Migrate()
	datafixtures.Load()
}

func main() {
	router.Routes()
}
