package main

import (
	"github.com/agnivade/levenshtein"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func fitFunction(attackUrl string, injectedBodyString string, xss string) float64 {

	var fitFunc float64
	var sigma float64 = 0
	var normalBodyString string
	var filtVariable float64

	resp, err := request(attackUrl)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		normalBodyString = string(bodyBytes)
	}

	levDist := float64(levenshtein.ComputeDistance(injectedBodyString, normalBodyString))
	lenNormalBodyString := float64(len(normalBodyString))
	lenInjectedBodyString := float64(len(injectedBodyString))

	xssLen := float64(len(xss))

	payloadLen := lenInjectedBodyString - lenNormalBodyString

	filteredChars := xssLen - payloadLen

	filtVariable = (xssLen - filteredChars) / xssLen
	pageDifference := levDist / lenNormalBodyString

	//fmt.Println("Xss len: ", xssLen)
	//fmt.Println("payloadLen: ", payloadLen)
	//fmt.Println("Filtered variable: ", filtVariable)
	//fmt.Println("Page difference: ", pageDifference)

	vr := VerifyReflection(injectedBodyString, xss)
	vd := VerifyDOM(injectedBodyString)

	if vr && vd {
		sigma = 100
		verifyChromedp(attackUrl + url.QueryEscape(xss))
	}

	fitFunc = sigma + 10*filtVariable + 5*pageDifference

	//fmt.Println("Sigma: ", sigma)
	//fmt.Println("Fit Function: ", fitFunc)
	//fmt.Println("-------------")

	return fitFunc

}
