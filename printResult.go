package main

import (
	"fmt"
	"github.com/mpvl/unique"
)
var xssFound []string

func printResult (){
	unique.Strings(&xssFound)
	for _, v := range xssFound {
		fmt.Println("Working XSS [+]: ", v)
	}
}
