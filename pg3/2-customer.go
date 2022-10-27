package pg3

import "fmt"

type customer struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	Phone     string
	Address   string
}

func (c *customer) Credentials() (string, string) {
	return c.Username, c.Password
}

func (c *customer) getAddress() string {
	return c.Address
}

func (c *customer) stdOutPrtln() {

	fmt.Printf("[customer.stdOutPrtln]\n")
	fmt.Printf("%s\n", c.FirstName)
	fmt.Printf("%s\n", c.LastName)
	fmt.Printf("%s\n", c.Username)
	fmt.Printf("%s\n", c.Password)
	fmt.Printf("%s\n", c.Email)
	fmt.Printf("%s\n", c.Phone)
	fmt.Printf("%s\n", c.Address)
}

func newCustomer(fn, ln, un, pw, em, ph, add string) *customer {
	return &customer{

		FirstName: fn,
		LastName:  ln,
		Username:  un,
		Password:  pw,
		Email:     em,
		Phone:     ph,
		Address:   add,
	}
}

func RunCustomer() {
	fmt.Printf("[Customer]\n")

	customer1 := newCustomer("Micheal", "Jordan", "MJ2020", "1234567", "MJ2020@gmail.com", "12345678", "18227 Capstan Greens Road Cornelius, NC 28031.")

	customer1.stdOutPrtln()

	username, password := customer1.Credentials()

	fmt.Printf("[RunCustomer] Username: %s Password: %s\n", username, password)

	address := customer1.getAddress()

	fmt.Printf("[RunCustomer] Address: %s \n", address)

}
