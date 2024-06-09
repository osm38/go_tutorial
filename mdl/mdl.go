package mdl

import "sample/logging"
import "golang.org/x/xerrors"

func Hello() {
	logger := logging.GetLogger()
	logger.Info("hello")
	err := xerrors.New("error!!!")
	logger.Eerror(err)
}
