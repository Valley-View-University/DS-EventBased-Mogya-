package logging

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../utility"
)

const loggerEndPoint = "mogya-dS/"

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r*http.Request){
			start:= time.Now()
			inner.ServeHTTP(w, r)

			timeTaken:= time.Since(start)
			timeTakenInNanoSeconds:= int64(timeTaken / time.Nanosecond)

			logMessage:= utility.APILogMessage{
				r.Method,
				r.RequestURI,
				name,
				timeTakenInNanoSeconds,
				"logs",
			}
			message := fmt.Sprintf(
				"%s\t%s\t%s\t%s\t%s",
				time.Now(),
				logMessage.Method,
				logMessage.URI,
				logMessage.Name,
				timeTaken,
			)
			log.Print(message)
		})
}