package main

import (
	"fmt"
	"log"

	"github.com/stundzia/amazon-scraper/goexamples"
)

func main() {
	b, err := goexamples.Amazon("https://www.amazon.co.uk/dp/B0BDJ279KF", true)
	if err != nil {
		log.Fatal("Amazon job failure: ", err)
	}
	fmt.Println("Amazon response: ", string(b))

	b, err = goexamples.AmazonBestsellers("automotive", "de", 2, 82400031, true)
	if err != nil {
		log.Fatal("Amazon Bestsellers job failure: ", err)
	}
	fmt.Println("Amazon Bestsellers response: ", string(b))
}
