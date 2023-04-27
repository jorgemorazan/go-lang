package main

import "fmt"

var l int = 45

var (
	pokemonName   string = "Pikachu"
	pokemonNumber int    = 25
	pokemonRegion string = "Kanto"
)

func main() {
	var i int
	i = 42
	var j int = 27
	k := 32
	fmt.Println(i, j, k, l)
	fmt.Println(pokemonName, pokemonNumber, pokemonRegion)
}
