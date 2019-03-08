package main

import (
  "fmt"
  "math/big"
)

/*recursive factorial function*/
func factorial( n *big.Int, result *big.Int, one *big.Int) *big.Int {
  // cmp returns -1 if n < 1
  // returns 0 if n == 1
  // so here I check n <= 1
  if n.Cmp(one) < 0 || n.Cmp(one) == 0 {
    return result
  }
  result.Mul(result, n)
  n.Sub(n, one)
  return factorial(n, result, one)
}

func main() {

  var n int64
  fmt.Scanf("%d\n", &n)
  N := big.NewInt(n)

  accumulate := big.NewInt(1)
  var one big.Int
  one.SetInt64(1)

  result := factorial(N, accumulate, &one)

  fmt.Printf("%d\n", result)
}
