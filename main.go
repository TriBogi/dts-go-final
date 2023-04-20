package main

import (
	"DTS-GO-FINAL/database"
	"DTS-GO-FINAL/routers"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	log.Println("Env successfully loaded")
}

func main() {
	database.StartDB()
	routers.StartApp().Run(":" + viper.GetString("PORT"))
}
