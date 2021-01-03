package main

import "golang.org/x/tour/pic"

//import "fmt"
//https://tour.golang.org/moretypes/18

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy, dy)
	for idx, _ := range ret {
		ret[idx] = make([]uint8, dx, dx)
	}
	// fmt.Println(ret)
	for idx, el := range ret {
		for idy, _ := range el {
			ret[idx][idy] = uint8((idx ^ idy))
		}
	}
	//fmt.Println(ret)
	return ret[:]
}

func main() {
	pic.Show(Pic)
}
