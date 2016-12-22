package main

import "fmt"

type Node struct {
  data int
  next *Node;
}

//Trying out self-referential structures
func main()  {
  var head Node
  head.data = 2
  var tail Node
  tail.data = 3
  head.next = &tail
  fmt.Println((*(head.next)).data)
}
