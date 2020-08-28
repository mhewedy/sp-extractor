package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	initConfig()
	initTypeMappings()
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
		_, _ = fmt.Fprintln(color.Output, color.RedString(err.Error()))
	}

	for _, mi := range mis {
		lines, err := mi.BodyAsLines()
		if err != nil {
			_, _ = fmt.Fprintln(color.Output, color.GreenString("\nStored Proc Body:> %s", lines))
			continue
		}

		spName, err := getSpName(lines)
		if err != nil {
			_, _ = fmt.Fprintln(color.Output, color.GreenString("\nStored Proc Body:> %s", lines))
			continue
		}

		_, _ = fmt.Fprintln(color.Output, color.HiGreenString("\n[Found] Stored Proc Name:> %s", spName))
	}
}
