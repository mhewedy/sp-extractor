package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var srcDir string

func init() {
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

	srcDir = strings.TrimSpace(string(b))
}

func main() {

	if len(os.Args) <= 3 {
		fmt.Printf("usage: %s <delegate class> <delegate method> <args number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	class := os.Args[1]
	method := os.Args[2]
	argsNumber, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	mi, err := parseInput(class, method, argsNumber)
	if err != nil {
		log.Fatal(err)
	}

	mis, err := getDAOMethodInfo([]MethodInfo{mi})
	if err != nil && err != errEndOfHierarchy {
		log.Fatal(err)
	}

	for _, mi := range mis {
		lines, err := mi.BodyAsLines()
		if err != nil {
			fmt.Println("\nStored Proc Body:> ", lines)
			continue
		}

		spName, err := getSpName(lines)
		if err != nil {
			fmt.Println("\nStored Proc Body:> ", lines)
			continue
		}

		fmt.Println("\n[Found] Stored Proc Name:> " + spName)
	}
}
