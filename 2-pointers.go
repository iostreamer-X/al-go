package main

import "fmt"

func main()  {
  a := 2
  pa := &a
  ppa := &pa

  //This is just me playing
  fmt.Println(pa,ppa)
}
