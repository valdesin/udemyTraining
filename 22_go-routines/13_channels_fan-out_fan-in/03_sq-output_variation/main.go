package main

import (
  "fmt"
  "sync"
)
func main() {
  in := gen(2,3)

  c1 := sq(in)
  c2 := sq(in)

  for v := range marge(c1,c2) {
    fmt.Println(v)
  }
}

func gen(num ...int) <-chan int {
  out := make(chan int)
  go func() {
    for _, v := range num {
      out <- v
    }
    close(out)
  }()
  return out
}

func sq(c <-chan int) <-chan int {
  out := make(chan int)
  go func() {
    for v := range c {
      out <- v * v
    }
    close(out)
  }()
  return out
}

func marge(ch ...<-chan int ) <-chan int {
  out := make(chan int)
  var wg sync.WaitGroup

  output := func(c <-chan int) {
      for v := range c {
        out <- v
      }
      wg.Done()
  }

  wg.Add(len(ch))
  for _, c := range ch {
    go output(c)
  }

  go func() {
    wg.Wait()
    close(out)
  }()
  return out
}
