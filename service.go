package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var saveFileInfo  = save
var spellCheckHelper = spellCheck
var readFileInfo = read

func process(fileName string, fileContent string) FileInfo {
	log.Println("Processing file upload.")
	words := strings.Fields(string(fileContent))
	fileInfo := FileInfo{
		FileName:  fileName,
		WordCount: make(map[string]int),
	}
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	for _, word := range words {
		fileInfo.WordCount[reg.ReplaceAllString(strings.ToLower(word), "")]++
	}
	saveFileInfo(fileInfo)
	return fileInfo
}

func query(fileName string, trim string, checkSpell bool) (FileInfo, []string, error) {
	log.Println("Reading file information.")
	fileInfo := readFileInfo(fileName)
	words := make([]string,0)

	if trim!="" || checkSpell {
		for e := range fileInfo.WordCount {
			if trim!="" && strings.HasPrefix(e, strings.ToLower(trim)) {
				delete(fileInfo.WordCount, e)
				continue
			}
			if checkSpell {
				words = append(words, e)
			}
		}
	}

	var misspelledWords []string
	var spellCheckErr error

	if checkSpell {
		misspelledWords, spellCheckErr = spellCheckHelper(words)
	}

	return fileInfo,misspelledWords, spellCheckErr
}

func spellCheck(words []string) ([]string, error) {

	req, err := http.NewRequest("POST", "https://api.cognitive.microsoft.com/bing/v7.0/spellcheck/?mkt=en-US&mode=proof", bytes.NewBuffer(buildBingRequestBody(words)))

	req.Header.Set("Ocp-Apim-Subscription-Key", *SPELL_CHECK_SUBSCRIPTION_KEY)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode!=200 {
		return nil, errors.New("Error in spell check api.")
	}
	body, _ := ioutil.ReadAll(resp.Body)
 	var spellcheckResults SpellCheck
	json.Unmarshal(body, &spellcheckResults)
	misspelledWords := make([]string,0)

	for _, flagToken := range spellcheckResults.FlaggedTokens {
		misspelledWords = append(misspelledWords, flagToken.Token)
	}

	return misspelledWords, nil
}

func buildBingRequestBody(words []string) []byte {
	body := "Text="
	for _,word := range words {
		body = body + word + ", "
	}
	return []byte(body)
}