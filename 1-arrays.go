package main

import "fmt"

func main(){
  array := []int {1,2,5,66}
  length,average := get_length_average(array)
  fmt.Printf("%d %f\n",length,average)
}

func get_length_average(array []int) (int,float32) {
  length := len(array)
  sum := get_sum(array)

  //This part is weird to me, I have to cast both the variables
  average := float32(sum)/float32(length)

  //This is nice, two results
  return length, average
}

func get_sum(array []int) int {
  sum := 0
  count := len(array)
  for index := 0; index < count; index++ {
    sum += array[index]
  }

  return sum
}
