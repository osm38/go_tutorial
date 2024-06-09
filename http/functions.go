package http

import (
	"net/http"
	"sample/logging"
)

var logger = logging.GetLogger()

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!")
	_, err := w.Write(hello)
	if err != nil {
		logger.Eerror(err)
	}
}
