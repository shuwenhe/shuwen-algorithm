package main

import(
	"fmt"
)

func main()  {
	gcd := gcd(18,27)
	fmt.Println("gcd = ",gcd)
}

func gcd(p,q int) int {
	if q == 0{
		return p
	}
	r := p % q
	return gcd(q,r)
}