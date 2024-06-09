package main

import (
	"fmt"
	http "net/http"
	"os"
	shttp "sample/http"
	"sample/logging"
)

var logger = logging.GetLogger()

func main() {
	http.HandleFunc("/hello", shttp.HelloHandler)
	fmt.Println("Server Start Up........")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Eerror(err)
		os.Exit(1)
	}
}
