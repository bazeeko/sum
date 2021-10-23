package main

import (
	"fmt"
	"github.com/bazeeko/utils"
)

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("The sum of the numbers is %v\n", utils.Sum(a, b))
}
