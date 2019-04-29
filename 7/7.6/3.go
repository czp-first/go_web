package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}

	b, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}

	c, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	d, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	e, err := strconv.Itoa("1023")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a, b, c, d, e)
}
