package main

import "fmt"

const PQ_SIZE = 10

type Heap struct{
  size int
  queue [PQ_SIZE]int
}

func swap(heap *Heap, index_1,index_2 int)  {
  temp := heap.queue[index_1]
  heap.queue[index_1] = heap.queue[index_2]
  heap.queue[index_2] = temp
}

//Some +1s and -1s to adjust for the index starting from 0
func get_children(index int) (int,int) {
  return 2*index + 1,2*index +2
}

func get_parent(index int) int {
  if(index == 0){
    return -1
  }

  return (index-1)/2
}

func insert(heap *Heap, item int)  {
  if(heap.size == PQ_SIZE){
    fmt.Println("Heap is full!")
    return
  }

  //Miss the old "heap.queue[heap.size++] = item" one liner
  heap.queue[heap.size] = item
  heap.size++

  bubble_up(heap, heap.size-1)
}

func bubble_up(heap *Heap, position int)  {
  parent_position := get_parent(position)

  if(parent_position == -1){
    return
  }

  if(heap.queue[parent_position] > heap.queue[position]){
    swap(heap,position,parent_position)
    bubble_up(heap,parent_position)
  }
}

func extract_min(heap *Heap) int {
  min := heap.queue[0]
  heap.queue[0] = heap.queue[heap.size-1]
  heap.size--

  bubble_down(heap,0)
  fmt.Println(heap.queue)
  return min
}

func bubble_down(heap *Heap, position int)  {
  childpos_1,childpos_2 := get_children(position)

  //Checking if the item is question isn't at the bottom of the queue
  if(childpos_1 < heap.size-1){

    //I initialzed this 0 with and couldn't get the actual lowest from children
    swap_position := position

    if(heap.queue[childpos_1] < heap.queue[position]){
      swap_position = childpos_1
    }
    if(heap.queue[childpos_2] < heap.queue[swap_position]){
      swap_position = childpos_2
    }

    //This condition prevents stackoverflow when array has 2 same elements.
    if(swap_position != position){
      swap(heap,position,swap_position)
      bubble_down(heap,swap_position)
    }
  }

  return
}

func make_heap(heap *Heap, array []int, array_size int)  {
  for index := 0; index < array_size; index++ {
    insert(heap,array[index])
  }
}

func heap_sort(heap *Heap, array []int, array_size int)  {
  for index := 0; index < array_size; index++ {
    array[index] = extract_min(heap)
  }
}

func main()  {
  var heap Heap
  array := []int {9,8,6,6,5,4}
  make_heap(&heap,array,6)
  heap_sort(&heap,array,6)
  fmt.Println(array)
}
