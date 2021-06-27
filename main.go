package main

import (
	"github.com/Timofey21/GoXSSfuzz/pkg/geneticAlgorithm"
	"github.com/Timofey21/GoXSSfuzz/pkg/read"
)

func main() {
	attackUrl := read.Args()
	geneticAlgorithm.GA(attackUrl, read.CSV())
}
