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

	_, err = mi.Next()
	if err != nil {
		log.Fatal(err)
	}

	/*mis, err := getDAOMethodInfo(mi)
	if err != errEndOfHierarchy {
		log.Fatal("EXIT", err)
	}

	for _, mi := range mis {
		lines, err := mi.bodyAsLines()
		if err != nil {
			log.Printf("error in %v\n", mi)
			continue
		}
		fmt.Println(getSpName(lines))
	}*/
}

func getDAOMethodInfo(mi MethodInfo) ([]MethodInfo, error) {

	arr, err := mi.Next()
	if err != nil {
		return nil, err
	}

	for _, mi := range arr {
		newArr, err := getDAOMethodInfo(mi)
		if err != nil {
			return nil, err
		}
		return newArr, nil
	}

	return nil, nil
}

func parseInput(input string) (MethodInfo, error) {

	// TODO implement

	return MethodInfo{
		class:      "VehiclesDelegate",
		method:     "listTrafficViolationsDetails",
		level:      levelDelegate,
		argsNumber: 1,
	}, nil
}

func getSpName(lines []string) (string, error) {
	fmt.Println("lines", lines)
	return "", nil
}
