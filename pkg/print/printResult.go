package print

import (
	"fmt"
	"github.com/Arafatk/glot"
	"github.com/mpvl/unique"
)
var XssFound []string
var FitFunc []float64

func PrintResult(){
	unique.Strings(&XssFound)
	for _, v := range XssFound {
		fmt.Println("Found XSS [+]: ", v)
	}
}

func Plot(){

	dimensions := 2
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	groupName := "Fit function"
	style := "lines"
	plot.SetTitle("Fit Function")
	plot.AddPointGroup(groupName, style, FitFunc)
	//plot.SavePlot("fitFunc.png")
}