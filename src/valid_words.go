package gordle

import (
	_ "embed"
	"strings"
)

//go:embed word-list.txt
var wordListContent string

var validWords = strings.Split(wordListContent, "\n")
