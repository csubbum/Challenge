package main

type FileInfo struct {
	FileName  string
	WordCount map[string]int
}

type SpellCheck struct
{
	Type string `json:"_type"`
	FlaggedTokens []SpellOffset `json:"flaggedTokens"`
}


type SpellOffset struct
{
	Offset int `json:"offset"`
	Token string `json:"token"`
}
