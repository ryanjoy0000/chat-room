package errPkg

import (
	"log"
	"net/http"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Panic(err)
	}
}

func HandleHttpErr(err error, msg string, w http.ResponseWriter, statusCode int) {
	if err != nil {
		log.Println(msg, ":", err)
		http.Error(w, msg, statusCode)
		return
	}
}