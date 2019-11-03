package main

import (
	"reflect"
	"testing"
)

func TestSaveSuccess(t *testing.T) {
	wordCountMap := make(map[string]int)
	wordCountMap["ABC"] = 3
	wordCountMap["CDE"] = 3
	wordCountMap["FFF"] = 1
	wordCountMap["EEE"] = 1
	fileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	save(fileInfo)

	if !reflect.DeepEqual(fileInfos["test.txt"], fileInfo) {
		t.Error("File save failed.")
	}
	delete(fileInfos, "test.txt")

}

func TestReadSuccess(t *testing.T) {
	wordCountMap := make(map[string]int)
	wordCountMap["ABC"] = 3
	wordCountMap["CDE"] = 3
	wordCountMap["FFF"] = 1
	wordCountMap["EEE"] = 1
	fileInfo := FileInfo{FileName: "test.txt", WordCount: wordCountMap}
	fileInfos["test.txt"] = fileInfo
	readInfo := read("test.txt")
	if !reflect.DeepEqual(readInfo, fileInfo) {
		t.Error("File read failed.")
	}
	delete(fileInfos, "test.txt")
}

func TestReadFail(t *testing.T) {
	readInfo := read("test.txt")
	fileInfo := FileInfo{
		FileName:  "",
		WordCount: map[string]int{},
	}
	if !reflect.DeepEqual(readInfo, fileInfo) {
		t.Error("File read failed.")
	}
	delete(fileInfos, "test.txt")
}