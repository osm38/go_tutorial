package main

import (
	"context"
	"encoding/json"
	"fmt"
	http "net/http"
	"os"
	shttp "sample/http"
	"sample/logging"
	"sample/model"

	"sample/aws"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var logger = logging.GetDefaultLogger()

func main() {
	client := aws.KinesisClient()

	user := &model.User{
		ID:   1,
		Name: "taro",
		Age:  50,
	}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
	pk := "dummy"
	stream := "TestStream"

	record := &kinesis.PutRecordInput{
		Data:         data,
		PartitionKey: &pk,
		StreamName:   &stream,
	}
	output, err := client.PutRecord(context.Background(), record, func(o *kinesis.Options) { o.Region = "ap-northeast-1" })
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
	}
	logger.Info(fmt.Sprintf("%+v", output))
}

func run() {
	http.HandleFunc("/hello", shttp.HelloHandler)
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
