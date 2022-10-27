package main

import (
	"fmt"
	"foodpanda-go-wk2/pg3"
	"os"
	"strconv"
)

var foos map[int]map[int]func()




func main() {


	foos[3][1] = pg3.Factorial

	args := os.Args
	args_l := len(args)
	if args_l == 1 {

		fmt.Println("No page given.")

		return
	}

	var page int
	if args_l == 2 {
		page, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("No page no given.")
			return
		}

	}

	fmt.Printf("Page %d\n", page)

	var activity int

	if args_l == 3 {
		activity, err = strconv.Atoi(args[2])

		if err != nil {
			fmt.Println("No activity no given.")
			return
		}
	}

}
