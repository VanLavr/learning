package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	p := new(Person)
	p.Name = "Ivan"
	p.Age = "20"
	p.Jb.Title = "Yan"
	p.Jb.Departament = "IT"
	fmt.Println("structure:")
	p.ShowInfo()

	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("\n\njson data (1st structure): %s\n\n", j)

	var p2 interface{}
	json.Unmarshal(j, &p2)
	fmt.Printf("json data (1st structure parsed to a 1st map): %v\n", p2)
	fmt.Printf("\n")

	m := map[string]string{
		"Name" : "Vasilii",
		"Age" : "21",
		"Job" : "DevOps",
	}

	fmt.Println("2nd map:")
	for k, v := range m {
		fmt.Printf("%s : %s\n", k, v)
	}

	j2, err2 := json.Marshal(m)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Printf("\n\njson data (from 2nd map): %s\n\n", j2)

	WriteBytes("person1.json", j)
}

type Job struct {
	Title 		string `json:"-	"`
	Departament string `json:"DEPARTAMENT"`
}

type Person struct {
	Name string `json:"FULLNAME"`
	Age  string `json:"AGE"`
	Jb   Job    `json:"JOB"`
}

func (p *Person) ShowInfo() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %s\n", p.Age)
	fmt.Printf("Company: %s\n", p.Jb.Title)
	fmt.Printf("Deptmnt: %s\n", p.Jb.Departament)
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

	fmt.Printf("Wrote %d bytes...\n", size)
}