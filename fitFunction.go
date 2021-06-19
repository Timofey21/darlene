package main

import (
	"github.com/agnivade/levenshtein"
	"io/ioutil"
	"log"
	"net/http"
)

func fitFunction(attackUrl string, injectedBodyString string, client *http.Client, xss string) (fitFunc float64) {

	var sigma float64 = 0
	var normalBodyString string

	resp, err := client.Get(attackUrl)
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


	//fmt.Println(normalBodyString)

	levDist := float64(levenshtein.ComputeDistance(injectedBodyString, normalBodyString))
	lenNormalBodyString := float64(len(normalBodyString))
	lenInjectedBodyString := float64(len(injectedBodyString))

	xssLen := float64(len(xss))

	payloadLen := lenInjectedBodyString - lenNormalBodyString

	filteredChars := xssLen - payloadLen
	pageDifference := levDist / lenNormalBodyString

	//fmt.Println("-------------")
	//fmt.Println("Filtered chars: ", filteredChars)
	//fmt.Println("Page difference: ", pageDifference)

	vr := VerifyReflection(injectedBodyString, xss)
	vd := VerifyDOM(injectedBodyString)

	if vr && vd {

		sigma = 1
		//verifyChromedp(attackUrl + url.QueryEscape(xss))
		//fmt.Println("XSS Found [+]		" + xssVector)
	}

	fitFunc = sigma + 4 * filteredChars + 2 * pageDifference


	//fmt.Println("Sigma: ", sigma)
	//fmt.Println("Fit Function: ", fitFunc)

	return

}