package main

import (
	"fmt"
	"log"
	"strings"
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

	/*mis, err := mi.Next()
	if err != nil {
		log.Fatal(err)
	}

	mis, err = mis[0].Next()
	if err != nil {
		log.Fatal(err)
	}

	mis, err = mis[0].Next()
	if err != nil {
		log.Fatal(err)
	}*/

	mis, err := getDAOMethodInfo([]MethodInfo{mi})
	if err != nil && err != errEndOfHierarchy {
		log.Fatal("EXIT", err)
	}

	for _, mi := range mis {
		lines, err := mi.BodyAsLines()
		if err != nil {
			log.Printf("error in %v\n", mi)
			continue
		}
		fmt.Println(getSpName(lines))
	}
}

func getDAOMethodInfo(mis []MethodInfo) ([]MethodInfo, error) {

	for _, mi := range mis {

		nextMis, err := mi.Next()
		if err != nil {
			if err == errEndOfHierarchy {
				return []MethodInfo{mi}, err
			}
			return nil, err
		}
		return getDAOMethodInfo(nextMis)

	}
	return nil, nil
}

func parseInput(input string) (MethodInfo, error) {

	// TODO implement

	return MethodInfo{
		class: "VehiclesDelegate",
		//method:     "listTrafficViolationsDetails",
		method:     "lookupLimitedVehiclePlateTypes",
		level:      levelDelegate,
		argsNumber: 1,
	}, nil
}

func getSpName(lines []string) (string, error) {
	fmt.Println("lines", strings.Join(lines, "\n"))
	return "", nil
}
