package log

import (
	"fmt"
	"time"
)

func Info(msg string) {
	fmt.Printf("[%v] %v\n", timestamp(), msg)
}

func Error(msg string) {
	fmt.Printf("[%v] ERROR! %v", timestamp(), msg)
}

func timestamp() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.9999")
}
