package main

import (
	"flag"
	"fmt"
	"time"
)

type Person struct {
	name string
	born time.Time
}

func (p Person) String() string {
	return fmt.Sprintf("Hello, my name is %s and I was born on %s", p.name, p.born)
}

func (p *Person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var person Person

	flag.Var(&person, "name", "the name of the person")

	flag.Parse()

	fmt.Println(person)
}
