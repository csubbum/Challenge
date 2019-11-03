package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUploadFile(t *testing.T) {
	srv := startServer()
	time.Sleep(2 * time.Second)
	bb := &bytes.Buffer{}

	writer := multipart.NewWriter(bb)
	part, err := writer.CreateFormFile("myFile", "test.txt")
	if err != nil {
		t.Error("Error creating multipart file")
	}
	_, err = io.Copy(part, strings.NewReader("Java red blue bluegreen orange. orange red orange yatpp samle"))
	if err != nil {
		t.Error("Error creating multipart file")
	}
	writer.Close()

	resp, respbody, err := callApi(bb, writer.FormDataContentType(), "http://localhost:8080/file/upload", "POST")

	if err != nil || resp.StatusCode != 200 {
		t.Error("Error uploading endpoint failed.")
	}
	wordCountMap := make(map[string]int)
	wordCountMap["java"] = 1
	wordCountMap["red"] = 2
	wordCountMap["blue"] = 1
	wordCountMap["bluegreen"] = 1
	wordCountMap["orange"] = 3
	wordCountMap["yatpp"] = 1
	wordCountMap["samle"] = 1
	expectedFileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}

	fileInfo := FileInfo{}
	json.Unmarshal(respbody, &fileInfo)

	if !reflect.DeepEqual(expectedFileInfo, fileInfo) {
		fmt.Println(expectedFileInfo)
		fmt.Println(fileInfo)
		t.Error("Failed in uploading the data")
	}

	// Query
	resp, respbody, err = callApi(&bytes.Buffer{}, "", "http://localhost:8080/file/query?fileName=test.txt&spellCheck=true&trim=blue", "GET")

	wordCountMap = make(map[string]int)
	wordCountMap["java"] = 1
	wordCountMap["red"] = 2
	wordCountMap["orange"] = 3
	wordCountMap["yatpp"] = 1
	wordCountMap["samle"] = 1
	expectedFileInfo = FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	dto := FileInfoDto{
		FileInfo:        expectedFileInfo,
		MisspelledWord:  []string{"yatpp", "samle"},
		SpellCheckError: "",
	}

	resultFileInfoDto := FileInfoDto{}
	json.Unmarshal(respbody, &resultFileInfoDto)

	if !reflect.DeepEqual(dto, resultFileInfoDto) {
		fmt.Println(dto)
		fmt.Println(resultFileInfoDto)
		t.Error("Failed in querying data")
	}

	// Error Scenario for upload
	resp, respbody, err = callApi(&bytes.Buffer{}, writer.FormDataContentType(), "http://localhost:8080/file/upload", "POST")

	if  resp.StatusCode != 400 {
		t.Error("Expecting error status 400 for upload")
	}

	// Error Scenario for query
	resp, respbody, err = callApi(&bytes.Buffer{}, "", "http://localhost:8080/file/query?spellCheck=true&trim=blue", "GET")

	if  resp.StatusCode != 400 {
		t.Error("Expecting error status 400 for query")
	}

	srv.Close()
}

func callApi(body *bytes.Buffer, contentType string, url string, requestMethod string) (*http.Response, []byte, error) {
	httpRequest, _ := http.NewRequest(requestMethod, url, bytes.NewReader(body.Bytes()))
	httpRequest.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, nil, err
	}
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	resp.Body.Close()
	return resp, respbody, nil
}