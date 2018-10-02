package main

import (
	"math"
	"strings"
)

type position struct {
	x int
	y int
}

func PeacefulYard(yard []string, minDistance int) bool {
	var L, M, R *position
	for i, v := range yard {
		Li := strings.Index(v, "L")
		Mi := strings.Index(v, "M")
		Ri := strings.Index(v, "R")

		if Li > -1 {
			L = &position{x: i, y: Li}
		}
		if Mi > -1 {
			M = &position{x: i, y: Mi}
		}
		if Ri > -1 {
			R = &position{x: i, y: Ri}
		}

		if L != nil && M != nil && R != nil {
			break
		}
	}
	minDist := float64(minDistance)

	if L != nil && M != nil {
		xLM := (L.x - M.x) * (L.x - M.x)
		yLM := (L.y - M.y) * (L.y - M.y)
		distLM := math.Sqrt(float64(xLM + yLM))

		if distLM > minDist {
			return false
		}
	}

	if L != nil && R != nil {
		xLR := (L.x - R.x) * (L.x - R.x)
		yLR := (L.y - R.y) * (L.y - R.y)
		distLR := math.Sqrt(float64(xLR + yLR))

		if distLR > minDist {
			return false
		}
	}

	if M != nil && R != nil {

	}
	xMR := (M.x - R.x) * (M.x - R.x)
	yMR := (M.y - R.y) * (M.y - R.y)
	distMR := math.Sqrt(float64(xMR + yMR))

	if distMR > minDist {
		return false
	}

	return true
}

func maincat() {
	yard := []string{"------------",
		"---M--------",
		"------------",
		"------------",
		"-------R----",
		"------------"}
	PeacefulYard(yard, 10)
}
