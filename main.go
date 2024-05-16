package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/retinacodeworks/regis/internal/server"
	"github.com/retinacodeworks/regis/internal/storage"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}

	databaseUrl := viper.Get("db").(string)
	fmt.Println(databaseUrl)
	// Create our database connection
	pool, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	//
	//m, err := migrate.New("file:///db/migrations", databaseUrl)
	//m.Up()

	// Create our file store. For now, we'll just default to minio for development
	store := storage.New(storage.NewMinioAdapter())

	// Initialize our server
	srv := server.New(pool, store)

	if err := srv.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/regis/")
	viper.AddConfigPath("$HOME/.regis")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
