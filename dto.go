package main

type FileInfoDto struct {
	FileInfo FileInfo
	MisspelledWord []string
	SpellCheckError string
}