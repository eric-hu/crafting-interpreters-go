package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!dlrow olleH"))
	fmt.Println(morestrings.ReverseRunes("!ηδΈε₯½δ½ "))
	fmt.Println(morestrings.ReverseRunes("!π§βπ³"))
	fmt.Println(morestrings.ReverseRunes("!π³βπ§"))
}
