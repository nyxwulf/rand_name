package main

import (
	"flag"
	"fmt"

	rn "github.com/nyxwulf/rand_name"
)

var right_name string

func init() {
	flag.StringVar(&right_name, "name", "", "Name")
	flag.StringVar(&right_name, "n", "", "Name")
}

func main() {
	flag.Parse()

	if right_name != "" {
		fmt.Println(rn.GetRandomName(rn.RightName(right_name)))
		return
	}

	fmt.Println(rn.GetRandomName())
}
