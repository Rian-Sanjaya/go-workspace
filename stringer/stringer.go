// One of the most ubiquitous interfaces is Stringer defined by the fmt package.

// type Stringer interface {
//    String() string
// }

// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.

package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// exercise
// make the IPAddr type implement fmt.Stringer to print the address as a dotted quad
// for instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4"
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
