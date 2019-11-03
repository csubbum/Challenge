package main

import (
	"fmt"
	"reflect"
	"testing"
)


func TestProcessSuccess(t *testing.T) {
	wordCountMap := make(map[string]int)
	wordCountMap["abc"] = 3
	wordCountMap["cde"] = 3
	wordCountMap["fff"] = 1
	wordCountMap["eee"] = 1
	expectedFileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	var mockSaved bool
	saveFileInfo = func(info FileInfo) {
		// mock
		mockSaved = true
	}
	actualFileInfo := process("test.txt", "ABC. ABC, CDE CDE cde abc fff eee")
	if !reflect.DeepEqual(expectedFileInfo, actualFileInfo) {
		t.Error("Failed in processing the data")
		fmt.Println(expectedFileInfo)
		fmt.Println(actualFileInfo)
	}
	if !mockSaved {
		t.Error("Save mock did not happen.")
	}
}

func TestQueryNoTrimSuccess(t *testing.T) {
	wordCountMap := make(map[string]int)
	wordCountMap["abc"] = 3
	wordCountMap["cde"] = 3
	wordCountMap["fff"] = 1
	wordCountMap["eee"] = 1
	expectedFileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	var mockSaved bool
	readFileInfo = func(fileName string) FileInfo {
		// mock
		mockSaved = true
		return expectedFileInfo
	}
	actualFileInfo, _, _ := query("test.txt", "", false)

	if !reflect.DeepEqual(expectedFileInfo, actualFileInfo) {
		t.Error("Failed in querying the data")
		fmt.Println(expectedFileInfo)
		fmt.Println(actualFileInfo)
	}
	if !mockSaved {
		t.Error("Query mock did not happen.")
	}
}

func TestQueryTrimSuccess(t *testing.T) {
	wordCountMap := make(map[string]int)
	wordCountMap["abc"] = 3
	wordCountMap["cde"] = 3
	wordCountMap["fff"] = 1
	wordCountMap["eee"] = 1

	fileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	var mockQuery bool
	readFileInfo = func(fileName string) FileInfo {
		// mock
		mockQuery = true
		return fileInfo
	}

	misspelledWords := []string{"abc","cde", "eee"}

	var mockSpellCheck bool
	spellCheckHelper= func(words []string) ([]string, error){
		// mock
		mockSpellCheck = true
		return misspelledWords, nil
	}

	wordCountMapExpected := make(map[string]int)
	wordCountMapExpected["abc"] = 3
	wordCountMapExpected["cde"] = 3
	wordCountMapExpected["eee"] = 1
	expectedFileInfo:= FileInfo{FileName: "test.txt", WordCount: wordCountMapExpected}

	actualFileInfo, spellCheckResults, _ := query("test.txt", "fff", true)

	if !reflect.DeepEqual(misspelledWords,spellCheckResults) {
		t.Error("Spellcheck failed")
	}
	if !reflect.DeepEqual(expectedFileInfo, actualFileInfo) {
		t.Error("Failed in querying the data")
		fmt.Println(expectedFileInfo)
		fmt.Println(actualFileInfo)
	}
	if !mockQuery {
		t.Error("Query mock did not happen.")
	}
	if !mockSpellCheck {
		t.Error("Query spell check did not happen.")
	}
}

func TestSpellCheckFail(t *testing.T) {
	*SPELL_CHECK_SUBSCRIPTION_KEY = "errorcode"
	_, err := spellCheck([]string{"random","sampl","yat"})
	if err==nil {
		t.Error("Query spell check should not pass.")
	}
}

