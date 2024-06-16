package main

import (
	"fmt"
	http "net/http"
	"os"
	shttp "sample/http"
	"sample/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var logger = logging.GetDefaultLogger()

func main() {
	run()
}

func run() {
	http.HandleFunc("/hello", shttp.HelloHandler)
	http.HandleFunc("/kinesis", shttp.KinesisHandler)
	fmt.Println("Server Start Up........")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Eerror(err)
		os.Exit(1)
	}
}

func generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection error!!")
	}

	g.UseDB(db)
	users := g.GenerateModel("users")
	sessions := g.GenerateModel("sessions")

	g.ApplyBasic(users, sessions)

	g.Execute()
}
