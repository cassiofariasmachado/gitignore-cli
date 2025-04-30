package log

import (
	"fmt"
	"log"
)

var (
	debugMode *bool
)

func Configure(debug *bool) {
	debugMode = debug
}

func Debug(format string, params ...interface{}) {
	if *debugMode {
		log.Printf(fmt.Sprintf("%s\n", format), params...)
	}
}

func Fatal(format string, params ...interface{}) {
	log.Fatalf(fmt.Sprintf("%s\n", format), params...)
}

func Print(format string, params ...interface{}) {
	log.Printf(fmt.Sprintf("%s\n", format), params...)
}
