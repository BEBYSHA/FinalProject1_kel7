package main

import (
	"github.com/Bobby-P-dev/FinalProject1_kel7/initializers"
	"github.com/Bobby-P-dev/FinalProject1_kel7/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
