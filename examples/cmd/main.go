package main

import (
	"fmt"
	socialmedia "reminders/examples/social-media"
)

func main() {

	umar := socialmedia.Facebook{UserName: "Umar", Email: "ufakz"}

	fmt.Println(umar.Feed())
	fmt.Println(umar.Fame())
}
