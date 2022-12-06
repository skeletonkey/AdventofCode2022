package adventOfCode

import (
	"bufio"
	"os"
)

var data *os.File

func Cleanup() {
	data.Close()
}

func GetData(path string, err error) *bufio.Scanner {
	ReportError(err)
	if os.Getenv("TESTCASE") == "1" {
		data, err = os.Open(path + string(os.PathSeparator) + "data_test")
	} else {
		data, err = os.Open(path + string(os.PathSeparator) + "data")
	}
	ReportError(err)

	return bufio.NewScanner(data)
}

func ReportError(err error) {
	if err != nil {
		panic(err)
	}
}
