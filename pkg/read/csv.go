package read

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type XssAttack struct {
	id          int
	Attack      []string
	FitFunction float64
}

func newXssAttack(id int, attack []string, fitFunction float64) XssAttack {
	return XssAttack{
		id:          id,
		Attack:      attack,
		FitFunction: fitFunction,
	}
}


func CSV() (xssAttacks []XssAttack) {
	csvFile, err := os.Open("payloads/xssAttacks.csv")

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

	return xssAttacks
}
