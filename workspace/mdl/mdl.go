package mdl

import (
	"sample/logging"

	"golang.org/x/xerrors"
)

func Hello() {
	logger := logging.GetDefaultLogger()
	logger.Info("hello")
	err := xerrors.New("error!!!")
	logger.Eerror(err)
}
