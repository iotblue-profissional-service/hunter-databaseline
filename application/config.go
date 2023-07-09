package application

import "os"

var (
	AUTHORIZATION_METHOD = os.Getenv("AUTHORIZATION_METHOD")
	USERNAME             = os.Getenv("USERNAME")
	PASSWORD             = os.Getenv("PASSWORD")
	envAuthUsername      = "gis@olympiccity.prv"
	envAuthPassword      = "gis1234"
)
