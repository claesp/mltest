package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func cost(data [][]float32, w float32, b float32) float32 {
	res := float32(0.0)
	for _, pair := range data {
		val := pair[0]
		exp := pair[1]
		act := val*w + b
		dif := act - exp
		res = res + (dif * dif)
		//fmt.Printf("actual: %f, expected: %f, diff: %f\n", act, exp, dif)
	}

	return res / float32(len(data))
}

func main() {
	t := [][]float32{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}

	r := 5
	if len(os.Args) > 1 {
		r, _ = strconv.Atoi(os.Args[1])
	}

	w := rand.Float32() * 10.0
	b := rand.Float32() * 5.0
	eps := float32(1e-3)
	rate := float32(1e-3)

	for i := 0; i < r; i++ {
		c := cost(t, w, b)
		dcostw := (cost(t, w+eps, b) - c) / eps
		dcostb := (cost(t, w, b+eps) - c) / eps
		w = w - rate*dcostw
		b = b - rate*dcostb
		//fmt.Printf("result: %f\n", cost(t, w))
	}

	fmt.Println("---")
	fmt.Println(w)
	fmt.Println("---")

	for _, p := range t {
		input := int(p[0])
		guess := int(math.Ceil(float64(w)))
		correct := int(p[1])
		fmt.Printf("%d * %d = %d (%d)\n", input, guess, (input * guess), correct)
	}
}
