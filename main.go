package main

import (
	"XSSfuzz/pkg/geneticAlgorithm"
	"XSSfuzz/pkg/read"
)

func main() {

	attackUrl := read.Args()
	geneticAlgorithm.GA(attackUrl, read.CSV())

}
