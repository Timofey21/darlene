package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func sum(array []float64) (result float64) {
	result = 0
	for _, v := range array {
		result += v
	}
	return
}

func main() {

	attackUrl := flag.String("url", "", "url")
	flag.Parse()

	var fitFunc []float64

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

	csvFile, err := os.Open("xssAttacks.csv")

	if err != nil {
		panic(err)
	}
	defer func(csv *os.File) {
		err := csv.Close()
		if err != nil {

		}
	}(csvFile)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.Comma = ';'
	reader.LazyQuotes = true

	xssAttacks, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 3;i++ {
		fmt.Println(xssAttacks[i][0] + xssAttacks[i][1] + xssAttacks[i][2] + xssAttacks[i][3] + xssAttacks[i][4] + xssAttacks[i][5])
	}


	fmt.Println(xssAttacks)

	f, err := os.Open("NewPayloads.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		xss := sc.Text()
		xssVector := *attackUrl + url.QueryEscape(xss)

		count++

		//start := time.Now()
		//go verification(xssVector)
		//duration := time.Since(start)
		//fmt.Println(duration)

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

			fitFunc = append(fitFunc, fitFunction(*attackUrl, bodyString, client, xss))

			//vr := VerifyReflection(bodyString, xss)
			//vd := VerifyDOM(bodyString)
			//
			//if vr && vd {
			//	verifyChromedp(xssVector)
			//	//fmt.Println("XSS Found [+]		" + xssVector)
			//}
		}

		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}

	}

	avgFitFunction := sum(fitFunc) / float64(len(fitFunc))
	fmt.Println("-------------")
	fmt.Println("-------------")
	fmt.Println("Average fit function: ", avgFitFunction)

}
