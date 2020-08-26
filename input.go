package main

func parseInput(class, method string, argsNumber int) (MethodInfo, error) {

	return MethodInfo{
		class:      class,
		method:     method,
		level:      levelDelegate,
		argsNumber: argsNumber,
	}, nil
}
