package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

func CountDigits(n int64) int {
	count := 0
	for d := n; d > 0; {
		d = d / 10
		count++
	}
	return count
}

type Sieve []bool

func (s Sieve) Show() {

	digits := CountDigits(int64(len(s)))
	padding := "%" + fmt.Sprintf("%d", digits)

	for i := 1; i < len(s); i++ {

		prime := " "
		if s[i] {
			prime = "p"
		}
		fmt.Printf((padding + "d(%s)| "), i, prime)

		if i%10 == 0 {
			fmt.Println()
		}
	}
}

func GetSieve(n int64) Sieve {

	sv := Sieve(make([]bool, n+1))

	for i := 0; i < len(sv); i++ {
		if (i >= 3 && i%2 != 0) || i == 2 {
			sv[i] = true
		}
	}

	stop := math.Sqrt(float64(n))
	//начинаем с 3
	for p := int64(3); float64(p) < stop; {
		//только для нечетных чисел
		if sv[p] {
			for i := p * p; i <= n; i += 2 * p {
				sv[i] = false
			}
		}
		p += 2
	}

	return sv

}

var an = flag.Int64("num", 0, "")

func main() {
	flag.Parse()
	log.Printf("get sieve for n=%d\r\n", *an)
	GetSieve(*an).Show()
}
