package main

import (
	"fmt"
	"hexArch/internal/adapters/core/arithmetic"
)

func main() {
	/* var core ports.ArithmeticPort

	core = arithmetic.NewAdapter() */

	arithmeticAdapter := arithmetic.NewAdapter()

	result, err := arithmeticAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
