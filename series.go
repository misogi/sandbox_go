package main

import (
  "fmt"
  "sync"
)

func GenerateSeries(max int) <- chan int {
  c := make(chan int)
  go func() {
    defer close(c)
    for i := 0; i < max; i++ {
      c <- i
    }
  }()
  return c
}

func ReadSeries(wg *sync.WaitGroup, c <-chan int, num int) {
  defer wg.Done()
  for i := range c {
    fmt.Printf("read %d %d \n", num, i)
  }
}

func main() {
  c := GenerateSeries(10)

  var wg sync.WaitGroup

  wg.Add(2)

  go ReadSeries(&wg, c, 1)
  go ReadSeries(&wg, c, 2)

  wg.Wait()
}