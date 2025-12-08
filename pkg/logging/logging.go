package log

import (
	"clean-arsitektur/pkg"
	"fmt"
	"net/http"
	"time"
)

func Log(message, level string, r *http.Request) {
	ip := pkg.ClientIP(r)
	t := time.Now()

	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"
	var log string

	switch level {
	case "INFO":
		log = green + "INFO" + reset
	case "ERROR":
		log = red + "ERROR" + reset
	case "WARN":
		log = yellow + "WARN" + reset
	}

	fmt.Println("[ LOG ] ", log, "--- Client IP:[", ip, "] Time:[", t.Format("2006-01-02 15:04:05"), "] Message:[", message, "]")
}
