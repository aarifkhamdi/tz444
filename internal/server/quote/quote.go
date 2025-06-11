package quote

import "math/rand"

var quotes = []string{
	"example1",
	"example2",
	"example3",
	"example4",
	"example5",
	"example6",
	"example7",
	"example8",
	"example9",
	"example10",
	"example11",
	"example12",
	"example13",
	"example14",
	"example15",
	"example16",
	"example17",
}

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
