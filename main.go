package main

import (
	"github.com/Timofey21/darlene/pkg/geneticAlgorithm"
	"github.com/Timofey21/darlene/pkg/read"
)

func main() {
	attackUrl := read.Args()
	geneticAlgorithm.GA(attackUrl, read.CSV())
}
