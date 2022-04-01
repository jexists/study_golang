package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)

	factorial := 1
	for i := 1; 1 <= num; i++ {
		factorial *= num
		num--
	}
	fmt.Println(factorial)
}

/*
n = eval(input())
result = 1

for i in range(1,n+1):
	result *= i

print(result)
*/

/*
def fibo(n):
    if n == 0:
        return 1
    else:
        return n * fibo(n-1)

n = eval(input())
print(fibo(n))
*/

/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	fmt.Println(factorial(n))
}

func factorial(n int) (result int) {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
*/
