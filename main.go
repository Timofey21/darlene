package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

)


func VerifyReflection(body, payload string) bool {
	if strings.Contains(body, payload) {
		return true
	}
	return false
}

func main() {

	attackUrl := flag.String("url", "", "url")
	flag.Parse()

	var proxy = "http://127.0.0.1:8080"
	proxyURL, err := url.Parse(proxy)

	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	f, err := os.Open("xssPayloads.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	count := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		xss := sc.Text()
		xssVector:= *attackUrl + url.QueryEscape(xss)

		resp, err := client.Get(xssVector)
		if err != nil {
			log.Println(err)
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			if VerifyReflection(bodyString, xss){
				fmt.Println("Found!	" + xssVector)
				count += 1

			}
		}

		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}

	}

	fmt.Println(count)

}
