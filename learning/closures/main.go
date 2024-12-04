package main

import "fmt"

func main() {
	harryPotterAggregator := addSpace()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
}

func addSpace() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}
