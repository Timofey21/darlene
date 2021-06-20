package main

import (
	"log"
	"net/http"
	"net/url"
)

func request(urlStr string) (resp *http.Response, err error) {

	var transport *http.Transport

	if *prox {
		proxy := "http://127.0.0.1:8080"
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Println(err)
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
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
