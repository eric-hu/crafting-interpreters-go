package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!dlrow olleH"))
	fmt.Println(morestrings.ReverseRunes("!界世好你"))
	fmt.Println(morestrings.ReverseRunes("!🧑‍🍳"))
	fmt.Println(morestrings.ReverseRunes("!🍳‍🧑"))
}
