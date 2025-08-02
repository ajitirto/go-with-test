package main 

import "fmt" 

func Hello() string {
  return "Hello world"
}

func HelloCustom(name string) string {
  return "hello, " + name
}

func main () {
  fmt.Println(Hello())
}
