package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	consulapi "github.com/habibiefaried/consul-file-server/consul"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	folder := "consulfs/"

	if os.Getenv("FOLDER") != "" {
		folder = os.Getenv("FOLDER")
		if folder[len(folder)-1] != '/' {
			folder = folder + "/"
		}
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	keyName := r.FormValue("key")
	if keyName == "" {
		_, _ = io.WriteString(w, "Variable 'key' must be provided\n")
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	err = consulapi.Upload(folder+keyName, buf.Bytes())

	if err != nil {
		_, _ = io.WriteString(w, fmt.Sprintf("%v", err))
	} else {
		_, _ = io.WriteString(w, "File uploaded")
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := fmt.Sprintf("You are hitting %s path", vars["filename"])
	_, _ = fmt.Fprintf(w, str)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods("GET", "HEAD")
	router.HandleFunc("/{filename}", homeLink).Methods("GET", "HEAD")
	router.HandleFunc("/", UploadFile).Methods("POST")

	httpPort := "8081"
	if os.Getenv("PORT") != "" {
		httpPort = os.Getenv("PORT")
	}
	log.Println("Listening on port " + httpPort)
	log.Print(http.ListenAndServe(":"+httpPort, router))
}
