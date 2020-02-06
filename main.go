package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//var alex person
	//fmt.Println(alex)

	jim := person{
		firstName: "jim",
		lastName:  "keller",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 95051,
		},
	}
	//jimpointer := &jim
	//fmt.Println(&jim)
	jim.updateName("JimmyJo")
	fmt.Println(jim)
	//jim.print()

	name := "bull"
	namePointer := &name

	//fmt.Println(*&name)
	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namepointer *string) {
	fmt.Println(&namepointer)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
