package main

import (
	"math"
	"math/rand"
)

var charset = map[string]string{
	"numbers":      "0123456789",
	"alphabetic":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"alphanumeric": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

func main() {

}

// randomInt, create random int (unsigned) //
// input uint
// output uint
func randomInt(min, max uint) uint {
	return math.Floor(rand.Uint32()*(max-min+1)) + min
}

// getCharset, get character set (numbers, alphabetic or alphanumeric) //
// input string
// output string
func getCharset(name string) string {
	return charset[name]
}

func randomElement(data string) string {
	return data[randomInt(0, len(data)-1)]
}

// repeat, repeat .....
// input string
// input uint
func repeat(input string, count uint) {
	var result string

	for cnt := 0; cnt < count; cnt++ {
		result += input
	}
	return result
}
