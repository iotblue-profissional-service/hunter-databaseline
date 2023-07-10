package main

import (
	"databaselineservice/application"
	"databaselineservice/sdk/cervello"
)

func init() {
	cervello.Login()
}

func main() {
	// application.StartBaselineApp()
	application.StartApiApp()
}
