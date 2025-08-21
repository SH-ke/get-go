package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// Length 5, capacity 5
	dwarfs2 := append(dwarfs1, "Orcus")
	// Length 6,capacity 10
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	// Length 9,capacity 10

	for _, s := range [][]string{dwarfs1, dwarfs2, dwarfs3} {
		fmt.Printf("Length %d, capacity %d\n", len(s), cap(s))
	}
}
