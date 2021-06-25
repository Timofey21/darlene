package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type xssAttack struct {
	id          int
	attack      []string
	fitFunction float64
}

func newXssAttack(id int, attack []string, fitFunction float64) xssAttack {
	return xssAttack{
		id:          id,
		attack:      attack,
		fitFunction: fitFunction,
	}
}


func readCSV () (xssAttacks []xssAttack) {
	csvFile, err := os.Open("xssAttacks.csv")

	if err != nil {
		panic(err)
	}
	defer func(csv *os.File) {
		err := csv.Close()
		if err != nil {

		}
	}(csvFile)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true

	buffer, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range buffer {
		xssAttacks = append(xssAttacks, newXssAttack(i, v, 0))
	}

	return
}
