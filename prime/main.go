package main

// 素数判断
import(
	"fmt"
)

func main()  {
	n := 4
	b := isPrime(n)
	if b == true{
		fmt.Println("n = ",n,"is prime")	
	}else{
		fmt.Println("n = ",n,"is not prime")	
	}
}

func isPrime(n int) (b bool) {
	if (n < 2){
		return false
	} 

	for i := 2; i*i <= n; i++{
		if n % i == 0 {
			return false
		}
	}
	return true
}