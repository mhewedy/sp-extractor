package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const mappingsCommentPrefix = "#"

type typeMapping []string

var typeMappings typeMapping

var (
	__homeDir, _    = os.UserHomeDir()
	mappingFilePath = filepath.Join(__homeDir, ".sp-extractor.map")
)

func (tm typeMapping) findNext(class string) (string, bool) {

	for _, line := range typeMappings {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, mappingsCommentPrefix) && strings.Contains(line, class) {

			part := strings.Split(line, ":")

			if part[0] == class {
				return part[1], true
			}
			if part[1] == class {
				return part[2], true
			}
		}
	}
	return "", false
}

func initTypeMappings() {

	f, err := os.OpenFile(mappingFilePath, os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	typeMappings = strings.Split(strings.TrimSpace(string(b)), "\n")
}
