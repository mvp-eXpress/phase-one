package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/mvp-eXpress/phase-one/dao-grpc"
	"github.com/mvp-eXpress/phase-one/dao-grpc/db"
)

type Config struct {
	DB        string `envconfig:"DB_TYPE"`
	MONGO_URL string `envconfig:"MONGO_URL"`
}

var mr todo.IRepo
var s todo.IService

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.DB == "MONGO" {

		fmt.Printf("DB type is %v", cfg.DB)
		fmt.Printf("MONGO_URL type is %v", cfg.MONGO_URL)

		mr, err := db.NewMongoConn(cfg.MONGO_URL)

		if err != nil {
			log.Fatal(err)
		}
		s = todo.NewService(mr)
		log.Fatal(todo.ListenGRPC(s, 8080))

	} else {
		log.Fatalln("DB type is not defined")
	}

}
