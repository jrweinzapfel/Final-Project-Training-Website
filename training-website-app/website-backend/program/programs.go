package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Programs struct {
	Programs []Program
}

type Program struct {
	Title string
	Type  string
	Split string
	ID    uint
	Body  string
}

func getJSON() Programs {
	jsonFile, err := os.Open("programs.json")
	if err != nil {
		fmt.Println(err)
		return Programs{}
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return Programs{}
	}
	var programs Programs
	err = json.Unmarshal(byteValue, &programs)
	if err != nil {
		fmt.Println(err)
		return Programs{}
	}
	return programs
}

func processPrograms(programs Programs) {
	for _, program := range programs.Programs {
		processProgram(program)
	}
}

func processProgram(program Program) {
	fmt.Printf("Processing Program: %+v\n", program)
}

func main() {
	programs := getJSON()
	processPrograms(programs)
}
