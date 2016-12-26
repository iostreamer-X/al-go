package main

import "fmt"

func quicksort(array []int, low,high int)  {

  //Using multiple returns for concise code
  left_low,left_high,right_low,right_high := partition(array,low,high)

  if(left_low < left_high){
    quicksort(array,left_low,left_high)
  }

  if(right_low < right_high){
    quicksort(array,right_low,right_high)
  }
}

func partition(array []int, low,high int) (int,int,int,int) {
  var left_low,left_high,right_low,right_high int

  //A bit of randomizing for edge cases input such as descending order array
  swap(array,low,(low+high)/2)
  pivot := low
  first_opened := pivot+1
  last_closed := -1

  for index := pivot+1; index <= high; index++ {
    if array[index] < array[pivot] {
      swap(array, index, first_opened)
      last_closed = first_opened
      first_opened++
    }
  }

  /*
    Messed up a bit here by overthinking and not using pivot as it should have
    been. Rather, I used last_closed to set bounds for the two sub arrays.
    And that created awkward situations such as negative index, infinite recursion.

    Lesson to take away, don't complicate logic.
  */
  if(last_closed != -1){
    swap(array,last_closed,pivot)
    pivot = last_closed
  }
  left_low = low
  left_high = pivot-1
  right_low = pivot+1
  right_high = high

  return left_low,left_high,right_low,right_high

}

func swap(array []int, pos1,pos2 int)  {
  temp := array[pos1]
  array[pos1] = array[pos2]
  array[pos2] = temp
}

func main()  {
  array := []int{9,8,7,6,5,4}
  quicksort(array,0,5)
  fmt.Println(array)
}
