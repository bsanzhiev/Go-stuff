package main

import "fmt"

func main() {
  naturals := make(chan int)
  squares := make(chan int)

  // Counter
  go func() {
    for i := 1; i <= 100; i++ {
      naturals <- i
    }
    close(naturals)
  }()

  go func() {
    for x := range naturals {
      squares <- x * x
    }
    close(squares)
  }()

  for x := range squares {
    fmt.Println(x)
  }
}
