package utils

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func NewLog() (err error) {
	Log, err = zap.NewProduction()
	if err != nil {
		return
	}
	return
}
