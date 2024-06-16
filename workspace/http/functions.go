package http

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sample/aws"
	"sample/logging"
	"sample/model"
	"sample/util"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

var (
	logger    = logging.GetDefaultLogger()
	names     = []string{"Micky", "Aloisius", "Irwin", "Dave", "Bart", "Barry", "Ralph", "Mychal", "Morris", "Maverick", "Melvil", "Rupert", "Rusty", "Eben", "Elton", "Colin", "Anthony", "Gil", "Kiernan", "Spencer", "Derrell", "Mat", "Clifton", "Giles", "Sean", "Lanny", "Ephraim", "Kameron", "Alphonse", "Burt", "Tobin", "Johnny", "Jonathon", "Erick", "Levi", "Roman", "Vincent", "Farrell", "Henry", "Faron", "Nolan", "Zac", "Alexander", "Dee", "Jamison", "Intyre", "Hugo", "Tracy", "Aurelius", "Sammie"}
	cuusers   = []string{"Mavis", "Janis", "Khloe", "Tracy", "Valerie", "Allegra", "Lyric", "Freda", "Abbi", "Beatrix", "Valorie", "Julian", "Danna", "Arlette", "Brianne", "Sawyer", "Justine", "Muriel", "Macy", "Mee", "Kaylie", "Maya", "Mariel", "Jojo", "Brigid", "Danika", "Kandace", "Tamara", "Adalee", "Rachelle", "Reina", "Gianna", "Tatum", "Sharla", "Candith", "Jordyn", "Dianne", "Blythe", "Alexandrea", "Wynter", "Braelynn", "Loreen", "Makenzie", "Cora", "Kyleigh", "Princess", "Sky", "Erma", "Beverly", "Dodie"}
	passwords = []string{"", "M8eAwbXY", "P5xwmkA7", "U8tZKPCX", "Lz28EXaM", "C3fgUAJM", "q2V83TLX", "y2DBzVSd", "yH3tvudb", "Si9qkv3K", "C2gHqVFJ", "Qw6anTp4", "x6TEZfMk", "Ri59pMDX", "G5qpsw2D", "N7gachis", "M6hkPuGy", "eP2EY7HG", "kU9ZyLhH", "Us3akACB", "Qd72cpmB", "eB8AYFmD", "e3Q5dBnU", "m6Lh7bUk", "j9XNvp6u", "Da3TNFYf", "N3iZqwC4", "Y5k7wQPA", "Gi7EakXj", "i6WYxBrk", "j6HmNXwf", "c7Sa8BVi", "y7QEPTH8", "g5EJNFDB", "fX8NRi6S", "Ng8WbhLf", "Wz8VnB5q", "b4S5PjYG", "bQ9Mk8eF", "sU65J8tE", "k2YLZGA9", "i8DbqB7z", "aV5DtKFf", "p2G6uTBX", "pK6tNVEn", "H6hnJDTB", "P7dvDSKX", "Qw5zjgv3", "i8Bq6jat", "v5TLHxMc", "w5QBP9Du"}
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!")
	_, err := w.Write(hello)
	if err != nil {
		logger.Eerror(err)
	}
}

func KinesisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	client := aws.KinesisClient()

	user := &model.User{
		ID:              rand.Int31(),
		Name:            names[rand.Intn(50)],
		Age:             rand.Int31n(100),
		Password:        passwords[rand.Intn(50)],
		LastLogin:       util.RandomDateTime(),
		Deleted:         false,
		CreateUser:      cuusers[rand.Intn(50)],
		CreateTimestamp: util.RandomDateTime(),
		UpdateUser:      cuusers[rand.Intn(50)],
		UpdateTimestamp: time.Now(),
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
	output, err := client.PutRecord(context.Background(), record, func(o *kinesis.Options) {})
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info(fmt.Sprintf("%+v", output))

	w.WriteHeader(http.StatusOK)
}
