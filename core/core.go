package core

import (
	"github.com/go-logr/logr"
)

func Hello(logger logr.Logger) {
	logger.V(0).Info("Hello, World!")
	logger.V(1).Info("This is a verbose message")
	logger.V(2).Info("This is a very verbose message")
	logger.V(3).Info("This is an extremely verbose message")
}
