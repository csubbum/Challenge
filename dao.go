package main

import "log"

var fileInfos map[string]FileInfo = make(map[string]FileInfo)

func save(fileInfo FileInfo){
	log.Println("Entering file data save.")
	fileInfos[fileInfo.FileName] = fileInfo
	log.Println("Completed file data save.")
}

func read(fileName string) FileInfo{
	log.Println("Entering file data read.")
	info := fileInfos[fileName]
	fileInfo := FileInfo{
		FileName:  "",
		WordCount: map[string]int{},
	}
	fileInfo.FileName = info.FileName

	for word, count := range info.WordCount {
		fileInfo.WordCount[word]=count
	}

	log.Println("File data completed.")
	return fileInfo
}