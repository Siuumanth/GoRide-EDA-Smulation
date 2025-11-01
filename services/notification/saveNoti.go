package notification

import (
	"fmt"
	"os"
	"time"
)

var notifChan = make(chan string, 100)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "%s: PANIC: %s\n", time.Now().Format(time.RFC3339), r)
	}
}

func init() {
	go func() {
		defer recoverFromPanic()
		f, err := os.OpenFile("notifications.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("File opened successfully")
		}

		defer f.Close()

		for msg := range notifChan {
			logEntry := fmt.Sprintf("%s: %s\n", time.Now().Format(time.RFC3339), msg)
			f.WriteString(logEntry)
		}
	}()
}

func SaveNotification(msg string) {
	notifChan <- msg
}
