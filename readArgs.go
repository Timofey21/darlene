package main

import "flag"

var prox *bool

func readArgs() string {
	attackUrl := flag.String("url", "", "url")
	prox = flag.Bool("proxy", false, "proxy")
	flag.Parse()

	return *attackUrl
}
