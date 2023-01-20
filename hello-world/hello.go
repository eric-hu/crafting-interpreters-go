package main

import (
  "fmt"
  "example/user/hello/morestrings"
)

func main() {
  fmt.Println(morestrings.ReverseRunes("!dlrow olleH"))
  fmt.Println(morestrings.ReverseRunes("!界世好你"))
  fmt.Println(morestrings.ReverseRunes("!🧑‍🍳"))
  fmt.Println(morestrings.ReverseRunes("!🍳‍🧑"))
}

