package main

import "flag"

var prox *bool

func readArgs() (attackUrl string){
	attackUrl = *flag.String("url", "", "url")
	prox = flag.Bool("proxy", false, "proxy")
	flag.Parse()

	return
}
