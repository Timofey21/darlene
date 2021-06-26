package read

import "flag"

var Prox *bool

func Args() string {
	attackUrl := flag.String("url", "", "url")
	Prox = flag.Bool("proxy", false, "proxy")
	flag.Parse()

	return *attackUrl
}
