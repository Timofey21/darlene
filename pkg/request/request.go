package request

import (
	"github.com/Timofey21/darlene/pkg/read"
	"log"
	"net/http"
	"net/url"
)

func Request(urlStr string) (resp *http.Response, err error) {

	var transport *http.Transport

	if *read.Prox != "no proxy" {
		proxyURL, err := url.Parse(*read.Prox)
		if err != nil {
			log.Println(err)
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = &http.Transport{
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err = client.Get(urlStr)
	if err != nil {
		log.Println(err)
	}

	return

}
