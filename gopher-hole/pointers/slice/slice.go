package main

import (
	"fmt"
	"unsafe"
)

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto"}

	reclassify(&planets)
	fmt.Println(planets)

	// Prints [Mercury Venus arth Mars Jupiter Saturn Uranus Neptune]
	// get the len capacity of the slice
	fmt.Println(len(planets), cap(planets))
	// get the illlegal index pluto
	// 注意：这种方式不安全，不推荐在生产代码中使用
	ptr := &planets[0] // 获取底层数组的指针
	pluto := (*[9]string)(unsafe.Pointer(ptr))[8]
	fmt.Println("Pluto:", pluto)
}
