package main

import (
	"fmt"
	"time"

	"github.com/acsellers/round"
)

func main() {
	start := time.Now()
	fmt.Println(len(round.Compilation(5000)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(len(round.Compilation(10000)))
	fmt.Println(time.Since(start))
}
