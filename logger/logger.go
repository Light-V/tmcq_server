package logger

import (
	"log"
	"os"
	"sync"
)

var mu sync.Mutex
var logger *log.Logger

func GetLogger() *log.Logger {
	mu.Lock()
	defer mu.Unlock()
	if logger == nil {
		logger = log.New(os.Stdout, "[test]", log.Lshortfile|log.Ldate|log.Ltime)
	}
	return logger
}
