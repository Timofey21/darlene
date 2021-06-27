package read

import "flag"

var Prox *string

func Args() string {
	attackUrl := flag.String("url", "", "url")
	Prox = flag.String("proxy", "no proxy", "proxy")
	flag.Parse()

	return *attackUrl
}
