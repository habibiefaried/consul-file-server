package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	keyName := r.FormValue("key")
	if keyName == "" {
		_, _ = io.WriteString(w, "Variable 'key' must be provided\n")
		return
	}

	fileName := "/tmp/" + handler.Filename

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	maxSizeInKb := 512
	if os.Getenv("MAX_FILESIZE") != "" {
		i, err := strconv.Atoi(os.Getenv("MAX_FILESIZE"))
		if err != nil {
			fmt.Println(err)
		} else {
			maxSizeInKb = i
		}
	}

	if handler.Size/1024 <= 512 {
		_, _ = io.Copy(f, file)
		_, _ = io.WriteString(w, "File "+fileName+" is uploaded successfully\n")
	} else {
		_, _ = io.WriteString(w, fmt.Sprintf("File %v has %vKB, max %vKB\n", fileName, handler.Size/1024, maxSizeInKb))
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
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
