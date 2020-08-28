package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var srcDirConfig string

func initConfig() {
	homeDir, _ := os.UserHomeDir()
	f, err := os.Open(filepath.Join(homeDir, ".sp-extractor"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	srcDirConfig = strings.TrimSpace(string(b))
}
