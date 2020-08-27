package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var srcDir string

func init() {
	srcDir = "C:\\Users\\mhewedy\\workspace\\code\\tamm\\Tamm-Portal"
	//srcDir = "/Users/mhewedy/Downloads/Tamm-Portal"
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
