package main

func main() {
  ch := make(chan struct{})
  go func() {
    ch <- struct{}{}
  }()
  <-ch
}