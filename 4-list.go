/*
  A queue implementation using self-referential structures.
  This was a good first step into data strucures, learnt pointers,
  how Go makes it easy to use pointers, and all in all a good revision.
*/

package main

import "fmt"

type Node struct{
  data int
  next *Node
}

//Realized recursion is going to come very handy
func search(list *Node, item int) *Node {
  if(list == nil){
    return nil;
  }

  if(list.data == item){
    return list
  }
  return search(list.next,item)
}

/*
  Concept of call by reference came rushing back.
  NOTE: Had list and tail been single pointers(*Node),
  I would have been operating on a pointer to the tail/head, and
  when I would have needed to update tail to point to a new node,
  I might have done &tail = &new_node, and this would have been very wrong.
  &tail would give the address of the argument tail, not actual tail.
*/
func insert(list **Node,tail **Node, item int)  {
  var new_node Node
  new_node.data = item
  if(*list == nil){
    *list = &new_node
    *tail = &new_node
    return
  }
  (*tail).next = &new_node
  *tail = &new_node
}

func predecessor(list *Node, item int) *Node {

  if(list == nil || list.next == nil){
    return nil
  }

  next_node := *(list.next)
  if(next_node.data == item){
    return list
  }else{
    return predecessor(list.next,item)
  }
}

func delete(list **Node, item int) {
  doomed := search(*list,item)

  if(doomed != nil){
    pred := predecessor(*list,item)
    if(pred == nil){
      *list = (*doomed).next
      return 
    }
    (*pred).next = (*doomed).next
  }
}

func traverse(node *Node)  {
  fmt.Println(node.data)

  if(node.next == nil){
    return
  }
  traverse(node.next)
}

func main()  {
  var head,tail *Node;
  insert(&head,&tail,2)
  insert(&head,&tail,22)
  insert(&head,&tail,23)
  insert(&head,&tail,12)

  fmt.Println(search(head,22))
  delete(&head,22)
  traverse(head)
}
