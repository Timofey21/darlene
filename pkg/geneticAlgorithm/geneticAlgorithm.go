package geneticAlgorithm

import (
	"github.com/Timofey21/darlene/pkg/print"
	"github.com/Timofey21/darlene/pkg/read"
	"github.com/Timofey21/darlene/pkg/request"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
)

const randomCircles = 20
const epoch = 100

func crossover(xssAttacks []read.XssAttack) []read.XssAttack {

	for i := 0; i < randomCircles; i++ {
		a1 := rand.Intn(len(xssAttacks))
		a2 := rand.Intn(len(xssAttacks))
		module := rand.Intn(6)

		buff := xssAttacks[a1].Attack[module]
		xssAttacks[a1].Attack[module] = xssAttacks[a2].Attack[module]
		xssAttacks[a2].Attack[module] = buff
	}

	return xssAttacks
}

func mutation(xssAttacks []read.XssAttack) []read.XssAttack {
	for i := 0; i < randomCircles; i++ {

		idx := rand.Intn(len(xssAttacks))
		module := rand.Intn(6)

		xssAttacks[idx].Attack[module] = xssAttacks[idx].Attack[module] + ""

	}

	return xssAttacks
}

func sum(xssAttacks []read.XssAttack) (result float64) {
	result = 0
	for _, v := range xssAttacks {
		result += v.FitFunction
	}
	return
}



func GA(attackUrl string, xssAttacks []read.XssAttack) {

	for i := 0; i < epoch; i++ {

		for j := 0; j < len(xssAttacks); j++ {

			xss := xssAttacks[j].Attack[0] + xssAttacks[j].Attack[1] + xssAttacks[j].Attack[2] +
				xssAttacks[j].Attack[3] + xssAttacks[j].Attack[4] + xssAttacks[j].Attack[5]

			xssVector := attackUrl + url.QueryEscape(xss)

			resp, err := request.Request(xssVector)
			if err != nil {
				log.Println(err)
			}

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)

				xssAttacks[j].FitFunction = fitFunction(attackUrl, bodyString, xss)
			}

			err = resp.Body.Close()
			if err != nil {
				log.Println(err)
			}
		}

		sort.SliceStable(xssAttacks, func(i, j int) bool {
			return xssAttacks[i].FitFunction > xssAttacks[j].FitFunction
		})

		newPopulation := xssAttacks // selection

		newPopulation = crossover(newPopulation)
		newPopulation = mutation(newPopulation)

		xssAttacks = newPopulation

		//avgFitFunction := sum(xssAttacks) / float64(len(xssAttacks))
		//fmt.Println("Epoch: ", i, "FitFunction: ", avgFitFunction)
	}

	print.PrintResult()
}
