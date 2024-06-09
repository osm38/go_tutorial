package main

import (
	"fmt"
	http "net/http"
	"os"
	shttp "sample/http"
	"sample/logging"
	"sample/util"
)

var logger = logging.GetDefaultLogger()

func main() {
	conf := &util.DbConfig{}
	util.LoadConfig("db-config.yml", conf)
	// logger.Debugf("%+v\n", conf)
	logConf := map[string]util.LogConfig{}
	util.LoadConfig("log-config.yml", logConf)
	logger.Debugf("%+v\n", logConf)

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
