package main

import (
	"fmt"
	"os"
//	"io"
	"ioutil"
	"bufio"
)

func ReadF(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}

	return fmt.Sprintf(string(content))
}

func ReadByLines(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ret, sep string
	sep = " "
	for scanner.Scan() {
		ret += fmt.Sprintf(scanner.Text())
		ret += sep
	}
	return ret
}

func main() {
	a := ReadF("example1.txt")
	fmt.Println(a)
	fmt.Println()

	b := ReadByLines("example1.txt")
	fmt.Println(b)
}