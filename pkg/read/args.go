package read

import (
	"flag"
	"fmt"
	"os"
)

var Prox *string

func Args() string {
	attackUrl := flag.String("url", "", "url")
	Prox = flag.String("proxy", "no proxy", "proxy")
	flag.Parse()

	if *attackUrl == ""{
		fmt.Println("Please add target url")
		os.Exit(0)
	}

	return *attackUrl
}
