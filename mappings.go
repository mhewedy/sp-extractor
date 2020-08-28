package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type typeMapping []string

var typeMappings typeMapping

func (tm typeMapping) findNext(class string) (string, bool) {

	for _, line := range typeMappings {
		if strings.Contains(line, class) {

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
	homeDir, _ := os.UserHomeDir()
	f, err := os.OpenFile(filepath.Join(homeDir, ".sp-extractor.map"), os.O_CREATE|os.O_RDONLY, 0600)
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
