//At this stage I read about Go directory strucure and importing rules

package main

import (
  "fmt"
  "./queue"
)

/*
  This one always made me a little fuzzy. The recursion here is a bit more
  than in previous programs and I wasn't comfortable with it, until now, almost.
*/
func mergesort(array []int, low,high int)  {
  if(low < high){
    middle := (low+high)/2
    mergesort(array,low,middle)
    mergesort(array,middle+1,high)
    merge(array,low,middle,high)
  }
}

func merge(array []int, low,middle,high int)  {
  var buffer1,buffer2,buffer1_tail,buffer2_tail *queue.Node

  //messed up a bit here by initializing this with 0
  counter := low

  for buffer1_index := low; buffer1_index <= middle; buffer1_index++ {
    queue.Insert(&buffer1,&buffer1_tail,array[buffer1_index])
  }

  for buffer2_index := middle+1; buffer2_index <= high; buffer2_index++ {
    queue.Insert(&buffer2,&buffer2_tail,array[buffer2_index])
  }

  for !(queue.Is_empty(buffer1) || queue.Is_empty(buffer2)) {
    if(queue.Peek(buffer1) < queue.Peek(buffer2)){
      array[counter] = queue.Dequeue(&buffer1)
    }else {
      array[counter] = queue.Dequeue(&buffer2)
    }

    counter++
  }

  for !queue.Is_empty(buffer1){
    array[counter] = queue.Dequeue(&buffer1)
    counter++
  }

  for !queue.Is_empty(buffer2){
    array[counter] = queue.Dequeue(&buffer2)
    counter++
  }
}

func main()  {
  array := []int {9,8,6,6,5,4}
  mergesort(array,0,5)
  fmt.Println(array)
}
