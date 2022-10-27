package pg7

import (
	"fmt"
	"reflect"
)

/*

Create a function inspect that would print out the type, kind and number of fields in the
struct.
The customer data are
1. First Name : “John”
2. Last Name: “Wick”
3. ID: 123123123
4. Invoice Total : 10000

*/

type customer struct {
	FirstName    string
	LastName     string
	ID           int
	InvoiceTotal float64
}

func inspect(i interface{}) {

	r := reflect.TypeOf(i)
	// k := r.Kind()
	nF := r.NumField()
	n := r.Name()

	v := reflect.ValueOf(i).Type()
	fmt.Printf("[inspect] Type: %s Number of Fields: %d StructName: %s\n", v, nF, n)
}

func RunStructReflection() {
	i := customer{
		FirstName:    "John",
		LastName:     "Wick",
		ID:           123123123,
		InvoiceTotal: 10000,
	}
	inspect(i)
}
