package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered file upload.")

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error uploading the File")
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error uploading the File")
		return
	}
	defer file.Close()

	fileBuffer, _ := ioutil.ReadAll(file)

	fileInfo := process(handler.Filename, string(fileBuffer))

	json.NewEncoder(w).Encode(fileInfo)
	log.Println("Completed file upload.")
}

func get(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering query data get")
	fileName := r.URL.Query()["fileName"]
	if fileName==nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "File name is not provided")
		return
	}
	trimQuery := r.URL.Query()["trim"]
	var trim string
	if trimQuery!=nil {
		trim = trimQuery[0]
	}

	spellCheckQuery := r.URL.Query()["spellCheck"]
	var spellCheck bool
	if spellCheckQuery!=nil {
		spellCheck = spellCheckQuery[0]=="true"
	}

	output, misspelledWord, spellcheckError := query(fileName[0],  trim, spellCheck)

	var errorMessage string
	if spellcheckError != nil {
		errorMessage = "Error during spell check."
	}

	fileInfoDto := FileInfoDto{
		FileInfo:       output,
		MisspelledWord: misspelledWord,
		SpellCheckError:errorMessage,
	}

	json.NewEncoder(w).Encode(fileInfoDto)
	log.Println("Completed query data get")
}

func startServer() *http.Server {
	log.Println("Server startup started.")
	http.HandleFunc("/file/upload", upload)
	http.HandleFunc("/file/query", get)
	srv := &http.Server{Addr: ":8080"}
	go func() {
		srv.ListenAndServe()
	}()
	log.Println("Server startup completed.")
	return srv
}
