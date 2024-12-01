package util

import (
	"bufio"
	"os"
	"path"
	"runtime"
)


func Read(name string) string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	b, err := os.ReadFile(path.Join(path.Dir(callingFile), name))
	HandleError(err)
	return string(b)
}

func ReadLines(name string) []string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	
	inputFile, err := os.Open(path.Join(path.Dir(callingFile), name))
	if err != nil {
		panic(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		HandleError(err)
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}
