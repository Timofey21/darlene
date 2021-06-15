package main

import (
	"bufio"
	"flag"
	"fmt"
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

//func VerifyDOM(s string) bool { //(body io.ReadCloser) bool {
//
//	body := ioutil.NopCloser(strings.NewReader(s)) // r type is io.ReadCloser
//	defer body.Close()
//
//	// Load the HTML document
//	doc, err := goquery.NewDocumentFromReader(body)
//	check := false
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	// Find the review items
//	doc.Find(".dalfox").Each(func(i int, s *goquery.Selection) {
//		check = true
//	})
//	if !check {
//		doc.Find("dalfox").Each(func(i int, s *goquery.Selection) {
//			// For each item found, get the band and title
//			check = true
//		})
//	}
//	return check
//}

func main() {

	attackUrl := flag.String("url", "", "url")
	flag.Parse()

	//var proxy = "http://127.0.0.1:8080"
	//proxyURL, err := url.Parse(proxy)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}
	//
	//client := &http.Client{
	//	Transport: transport,
	//}

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

		verification(xssVector)


	//	resp, err := client.Get(xssVector)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//
	//	if resp.StatusCode == http.StatusOK {
	//		bodyBytes, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		bodyString := string(bodyBytes)
	//		if VerifyReflection(bodyString, xss){
	//			verification(xssVector)
	//		}
	//	}
	//
	//	err = resp.Body.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//
	}

	fmt.Println(count)


}
