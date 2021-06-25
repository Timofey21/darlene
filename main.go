package main

func main() {


	attackUrl := readArgs()
	println(attackUrl)
	GA(attackUrl, readCSV())

}
