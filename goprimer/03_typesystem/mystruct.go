package main

import (
	"fmt"
)

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

var bill user;

func main() {

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println(lisa.name);
	fmt.Println(lisa.email);
	fmt.Println(lisa.ext);
	fmt.Println(lisa.privileged);

}
