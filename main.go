package main

import (
	"XSSfuzz/pkg/GA"
	"XSSfuzz/pkg/read"
)

func main() {

	attackUrl := read.Args()
	GA.GA(attackUrl, read.CSV())

}
