package main

import "strings"

func parseInput(class, method string, argsNumber int) (MethodInfo, error) {

	return MethodInfo{
		module:     strings.TrimSuffix(class, "Delegate"),
		class:      class,
		method:     method,
		level:      levelDelegate,
		argsNumber: argsNumber,
	}, nil
}
