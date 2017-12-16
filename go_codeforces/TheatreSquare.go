package main

import (
	"fmt"
	"math/big"
)


func getNoSquares(m, n, a big.Int) big.Int {
	return big.Int.Mul(big.Int.Div(big.Int.Add(m, a - 1), a), big.Int.Div(big.Int.Add(n, a - 1), a))
}

func main() {
	var m, n, a big.Int
	fmt.Scanf("%d %d %d", &m, &n, &a)
	fmt.Println(getNoSquares(m, n, a))
}
