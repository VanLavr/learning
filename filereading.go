package main

import (
	"fmt"
	"os"
	"ioutil"
	"bufio"
	"log"
)

func ReadAllInfo(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}

	return fmt.Sprintf(string(content))
}

func WriteStringInfo(args string, path string) {
	file, _ := os.Create(path)
	defer file.Close()

	_, err := file.WriteString(args)
	if err != nil {
		log.Panicf("Failed to wrote: %v\n", err)
	}
}

func WriteBytes(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	size, err := file.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Wrote %d bytes...", size)
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
	// reading all
	a := ReadAllInfo("example1.txt")
	fmt.Println(a)
	fmt.Println()

	// reading by lines
	b := ReadByLines("example1.txt")
	fmt.Println(b)

	//writing...
	path := "C:\\Users\\Ivan\\GoLang\\learning\\example2.txt"
	WriteStringInfo("Wrote this info in example2.txt", path)

	a = ReadAllInfo("example1.txt")
	fmt.Println(a)
	fmt.Println()
	a = ReadAllInfo("example2.txt")
	fmt.Println(a)
	fmt.Println()

	s := []byte{1, 2, 3, 4, 5, 6, 7, 9}
	WriteBytes("example3", s)
}