package main

import (
	"fmt"
	"foodpanda-go-wk2/pg3"
	"foodpanda-go-wk2/pg4"
	"foodpanda-go-wk2/pg5"
	"foodpanda-go-wk2/pg6"
	"foodpanda-go-wk2/pg7"
	"os"
	"strconv"
)

type category struct {
	page     int
	activity int
}

func main() {
	fmt.Println("main")

	foos := make(map[int]map[int]func())

	foos[3] = make(map[int]func(), 0)
	foos[3][1] = pg3.Factorial
	foos[3][2] = pg3.RunCustomer

	foos[4] = make(map[int]func(), 0)
	foos[4][1] = pg4.Games

	foos[5] = make(map[int]func(), 0)
	foos[5][1] = pg5.RunInterface

	foos[6] = make(map[int]func(), 0)
	foos[6][1] = pg6.RunBasicReflection

	foos[7] = make(map[int]func(), 0)
	foos[7][2] = pg7.RunStructReflection

	args := os.Args
	args_l := len(args)
	if args_l == 1 {

		fmt.Println("No page given.")

		return
	}

	var cat category

	if args_l == 2 {

		fmt.Println("No activity no given.")

	}

	if args_l == 3 {
		page_, err := strconv.Atoi(args[1])
		cat.page = page_
		if err != nil {
			fmt.Println("No valid page given.")
			return
		}
		activity_, err := strconv.Atoi(args[2])
		cat.activity = activity_
		if err != nil {
			fmt.Println("No valid activity given.")
			return
		}

		fmt.Printf("Page %d Activity %d\n", cat.page, cat.activity)

	}

	if foos[cat.page] != nil && foos[cat.page][cat.activity] != nil {
		foos[cat.page][cat.activity]()

	} else {

		fmt.Printf("No implementation for Page %d Activity %d \n", cat.page, cat.activity)
	}

}
