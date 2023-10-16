package main

import (
	"github.com/Bobby-P-dev/FinalProject1_kel7/initializers"
	router "github.com/Bobby-P-dev/FinalProject1_kel7/routes"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	router.Routes()
}
