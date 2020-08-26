package main

import (
	"errors"
	"regexp"
	"strings"
)

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

func getSpName(lines []string) (string, error) {
	for _, line := range lines {
		if strings.Contains(line, ".prepareCall") {
			p, err := regexp.Compile(`\.prepareCall\((.*)\);`)
			if err != nil {
				return "", err
			}
			m := p.FindAllStringSubmatch(line, -1)
			if m != nil {
				return strings.TrimSpace(m[0][1]), nil
			}
		}
	}
	return "", errors.New("cannot find procedure in DAO")
}
