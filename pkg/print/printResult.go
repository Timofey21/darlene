package print

import (
	"fmt"
	"github.com/mpvl/unique"
)
var XssFound []string

func PrintResult(){
	unique.Strings(&XssFound)
	for _, v := range XssFound {
		fmt.Println("Found XSS [+]: ", v)
	}
}
