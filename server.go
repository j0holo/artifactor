package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const MaxMemory = 10 << 20 // 20MB

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(MaxMemory)
		if err != nil {
			log.Println("Couldn't parse multipartform.")
			log.Println(err)
			// TODO: Maybe expand this later with various error messages that matches the error.
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		uploadedFile, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}
		defer uploadedFile.Close()

		log.Printf("Uploaded File: %+v\n", handler.Filename)
		log.Printf("File Size: %+v\n", handler.Size)
		log.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := ioutil.ReadAll(uploadedFile)
		if err != nil {
			log.Println("Could not read uploaded file, error:", err)
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = ioutil.WriteFile(fmt.Sprintf("artifact/%s", handler.Filename), fileBytes, os.FileMode(0660))
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}
		fmt.Fprint(w, http.StatusText(http.StatusAccepted))

	} else {
		fmt.Fprintf(w, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func main() {

	http.Handle("/artifact/", http.StripPrefix("/artifact/", http.FileServer(http.Dir("./artifact/"))))
	http.HandleFunc("/upload", upload)

	log.Println("Starting server on: http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
