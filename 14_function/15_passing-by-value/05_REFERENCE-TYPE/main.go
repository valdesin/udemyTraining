package main

import "fmt"

func main() {
  m := make(map[string]int)
  changeMe(m)
  fmt.Println(m["Vlad"])
}

func changeMe(m map[string]int) {
  m["Vlad"] = 14
}
