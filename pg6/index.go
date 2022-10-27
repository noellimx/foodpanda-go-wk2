package pg6

import (
	"fmt"
	"reflect"
)

func inspect(some ...interface{}) {

	for idx, intf := range some {

		t := reflect.TypeOf(intf)
		k := t.Kind()
		fmt.Printf("%d Type: %-15s Value: %s\n", idx, k, intf)
	}

}
func RunBasicReflection() {

	heteros := []interface{}{"This is a string", 12345, 1.2345, true}
	inspect(heteros...)
}
