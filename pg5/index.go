package pg5

import (
	"fmt"
	"reflect"
)

func printItemsAsStrings[U PrintableItem](stringables []U) {

	reflection := reflect.TypeOf(reflect.ValueOf(reflect.TypeOf(stringables).Elem()))

	fmt.Printf("[printItemAsStrings]Kind %s\n", reflection.Kind())
	for idx, s := range stringables {
		fmt.Printf("%-10d %s\n", idx+1, (s).AsString())
	}
}
func RunInterface() {

	fmt.Printf("[RunInterface]\n")
	b1 := newBook("Candle in the tomb", 10)
	b2 := newBookT("book_2", 20)
	bs := []*book{b1, b2}

	ca1 := newComputerAccessory("Razer BT Earpiece", 159)

	cas := []*computerAccessory{ca1}

	printItemsAsStrings(bs)
	printItemsAsStrings(cas)
}
