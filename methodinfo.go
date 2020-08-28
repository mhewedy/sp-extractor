package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var errEndOfHierarchy = errors.New("end of class hierarchy")

const newline = "\r\n"

const (
	levelDelegate level = iota
	levelBean
	levelDAO
)

const (
	levelDelegateSuffix = "Delegate"
	levelBeanSuffix     = "Bean"
	levelDAOSuffix      = "DAO"
)

const local = "Local"

type level uint8

func (c level) next() (level, error) {
	if c == levelDAO {
		return 0, errEndOfHierarchy
	}
	return c + 1, nil
}

func (c level) String() string {
	switch c {
	case levelDelegate:
		return levelDelegateSuffix
	case levelBean:
		return levelBeanSuffix
	case levelDAO:
		return levelDAOSuffix
	}
	return ""
}

func (c level) searchKeywords() []string {
	switch c {
	case levelDelegate:
		return []string{levelDelegateSuffix}
	case levelBean:
		return []string{levelBeanSuffix, local}
	case levelDAO:
		return []string{levelDAOSuffix}
	}
	return nil
}

// -----------

type MethodInfo struct {
	module     string
	class      string
	method     string
	level      level
	argsNumber int
	lines      []string
}

// ----
func (m *MethodInfo) Next() ([]MethodInfo, error) {

	lines, err := m.BodyAsLines()
	if err != nil {
		return nil, err
	}

	nextLevel, err := m.level.next()
	if err != nil {
		return nil, err
	}

	var mis []MethodInfo

	for _, line := range lines {
		info, found, err := m.findMethodInfo(line, nextLevel)
		if err != nil {
			return nil, err
		}
		if found {
			mis = append(mis, info)
		}
	}

	if len(mis) == 0 {
		return nil, fmt.Errorf("%s cannot find next classes in the call heiraricy", m.class)
	}

	return mis, nil
}

func (m *MethodInfo) findMethodInfo(line string, nextLevel level) (MethodInfo, bool, error) {
	//fmt.Println("Level:", nextLevel, "Line:", line)

	for _, keyword := range nextLevel.searchKeywords() {

		if strings.Contains(line, keyword) {

			p, err := regexp.Compile(fmt.Sprintf(`(?:=|return)(.*%s)[\.getInstance()]?.*?\.(.*?)\((.*)\);`, keyword))
			if err != nil {
				return MethodInfo{}, false, err
			}

			//fmt.Println("pattern", p)

			found := p.FindAllStringSubmatch(line, -1)
			if found != nil {

				class, ok := typeMappings.findNext(m.class)
				if !ok {
					class = m.module + nextLevel.String()
				}
				method := strings.TrimSpace(found[0][2])
				args := strings.TrimSpace(found[0][3])
				argsNumber := 0
				if len(args) > 0 {
					argsNumber = 1
				}
				argsNumber += strings.Count(args, ",")

				return MethodInfo{
					module:     m.module,
					class:      class,
					method:     method,
					argsNumber: argsNumber,
					level:      nextLevel,
				}, true, nil
			}
		}
	}

	return MethodInfo{}, false, nil
}

// ----

func (m *MethodInfo) BodyAsLines() ([]string, error) {
	if len(m.lines) > 0 {
		return m.lines, nil
	}

	classPath, err := findClassPath(m.class)
	if err != nil {
		return nil, err
	}

	contents, err := readAndNormalizeContents(classPath)
	if err != nil {
		return nil, err
	}

	methodIndex, err := findMethodIndex(contents, m.method, m.class, m.argsNumber)
	if err != nil {
		return nil, err
	}

	body, err := findBody(contents, methodIndex[1])
	m.lines = body
	return body, err
}

func findClassPath(class string) (string, error) {
	var classPath string
	err := filepath.Walk(srcDirConfig, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, fmt.Sprintf("%s%s%s", string(filepath.Separator), class, ".java")) {
			classPath = path
			return io.EOF
		}
		return nil
	})
	if err != nil && err != io.EOF {
		return "", err
	}
	if classPath == "" {
		return "", fmt.Errorf("cannot find file for class: %s", class)
	}
	return classPath, nil
}

func readAndNormalizeContents(classPath string) (string, error) {
	file, err := os.Open(classPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	// split content on semicolon (instead of newline)
	c := string(b)
	c = strings.ReplaceAll(c, newline, " ")
	c = strings.ReplaceAll(c, ";", ";"+newline)
	return c, nil
}

func findMethodIndex(contents, method, class string, argsNumber int) ([]int, error) {

	methodSigRegex := fmt.Sprintf("%s%s%s", `\s+`, method, `\s*\(`)
	for i := 0; i < argsNumber; i++ {
		methodSigRegex += ".*?," // ? for non greedy
	}
	if argsNumber > 0 {
		methodSigRegex = strings.TrimSuffix(methodSigRegex, ",")
	}
	methodSigRegex += `\)`

	p, err := regexp.Compile(methodSigRegex)
	if err != nil {
		return nil, err
	}

	methodSigIndex := p.FindStringIndex(contents)

	if len(methodSigIndex) == 0 {
		return nil, fmt.Errorf("cannot find method: %s with number of args %d in class %s",
			method, argsNumber, class)
	}

	fmt.Printf("Found: %s.%s with (%d) args\n", class, method, argsNumber)

	return methodSigIndex, nil
}

func findBody(content string, startIndex int) ([]string, error) {
	var body string

	var openCount, closeCount = 0, 0
	for i := startIndex; i < len(content); i++ {
		char := string(content[i])
		body += char

		if char == "{" {
			openCount += 1
		}
		if char == "}" {
			closeCount += 1
		}
		if openCount > 0 && openCount == closeCount {
			break
		}
	}
	return strings.Split(body, newline), nil
}
