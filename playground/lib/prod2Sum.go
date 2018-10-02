package main

import (
	"fmt"
	"math"
)

// Prod2Sum func
func Prod2Sum(a, b, c, d int) (rs [][]int) {
	ab := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
	cd := math.Pow(float64(c), 2) + math.Pow(float64(d), 2)
	n := ab * cd
	mainGuess := math.Round(math.Sqrt(n))

	var exist []int

Loop:
	for i := int(mainGuess); i > 0; i-- {
		guessSub := n - math.Pow(float64(i), 2)
		guessSub = math.Round(math.Sqrt(guessSub))
		sumGuess := math.Pow(float64(i), 2) + math.Pow(float64(guessSub), 2)
		if n == sumGuess {
			for _, e := range exist {
				if e == int(guessSub) {
					continue Loop
				}
			}
			exist = append(exist, i)
			intGuessSub := int(guessSub)
			if intGuessSub < i {
				rs = append(rs, []int{intGuessSub, i})
			} else {
				rs = append(rs, []int{i, intGuessSub})
			}
		}
	}
	return
}

func prod2SumMain() {
	rs := Prod2Sum(10, 11, 12, 13)
	fmt.Printf("%v", rs)
}
