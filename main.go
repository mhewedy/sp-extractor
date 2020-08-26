package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var srcDir string

func init() {
	//srcDir = "C:\\Users\\mhewedy\\workspace\\code\\tamm\\Tamm-Portal"
	srcDir = "/Users/mhewedy/Downloads/Tamm-Portal"
}

func main() {

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
			log.Printf("error in %v\n", mi)
			continue
		}

		spName, err := getSpName(lines)
		if err != nil {
			log.Printf("error in %v\n", mi)
			continue
		}

		fmt.Println("\nStored Proc:> " + spName)
	}
}
