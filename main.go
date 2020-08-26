package main

import (
	"fmt"
	"log"
)

var srcDir string

func init() {
	//srcDir = "C:\\Users\\mhewedy\\workspace\\code\\tamm\\Tamm-Portal"
	srcDir = "/Users/mhewedy/Downloads/Tamm-Portal"
}

func main() {

	input := "getVehiclesDelegate().lookupLimitedVehiclePlateTypes(getUserInfo())"

	mi, err := parseInput(input)
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

func parseInput(input string) (MethodInfo, error) {

	// TODO implement

	return MethodInfo{
		class: "IstemaraDelegate",
		//method:     "listTrafficViolationsDetails",
		method:     "listActiveAssignedPlates",
		level:      levelDelegate,
		argsNumber: 3,
	}, nil
}
