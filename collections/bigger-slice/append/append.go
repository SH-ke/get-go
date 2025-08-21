package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	// Prints [Ceres Pluto Haumea Makemake Eris Orcus]

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
	// Prints [Ceres Pluto Haumea akemake Eris Orcus Salacia Quaoar Sedna]
}
