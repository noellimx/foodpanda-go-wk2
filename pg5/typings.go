package pg5

import "fmt"

type book struct {
	Title string
	Price float64
}

func (b *book) AsString() string {

	return fmt.Sprintf("%s, $%0.5f", b.Title, b.Price)
}

func newBook(t string, p float64) *book {
	return &book{Title: t, Price: p}
}

func genNewT[U book | computerAccessory]() func(t string, p float64) *U {

	return func(t string, p float64) *U {

		return &U{
			Title: t,
			Price: p,
		}
	}

}

var newBookT = genNewT[book]()
var newComputerAccessory = genNewT[computerAccessory]()

type computerAccessory struct {
	Title string
	Price float64
}

func (ca *computerAccessory) AsString() string {

	return fmt.Sprintf("%s, $%0.5f", ca.Title, ca.Price)
}

type PrintableItem interface {
	AsString() string
}
