/*
  This is where things got interesting and heavy. And this is where I learnt
  the most about using recursion. And most importantly, handling arguments
  smartly.

  Here I heavily used Go's feature of letting me return multiple
  values. This worked like synergy, I return the item, as well as its address.
*/

package main

import "fmt"

type Tree struct{
  data int
  left *Tree
  right *Tree
  parent *Tree
}

func insert(tree **Tree,item int,parent *Tree)  {

  if(*tree == nil){
    new_tree := Tree{item,nil,nil,parent}
    *tree = &new_tree
    return
  }

  if(item < (*tree).data){
    insert(&(*tree).left,item,*tree)
  }else{
    insert(&(*tree).right,item,*tree)
  }
}

func traverse(root *Tree)  {
  if(root != nil){
    traverse(root.left)
    fmt.Println(root.data)
    traverse(root.right)
  }
}

func min(tree **Tree) (*Tree,**Tree)  {
  if((*tree).left == nil){
    return *tree,tree
  }
  return min(&(*tree).left)
}

func max(tree **Tree) (*Tree,**Tree)  {
  if((*tree).right == nil){
    return *tree,tree
  }
  return max(&(*tree).right)
}

func search(tree **Tree, item int) (*Tree,**Tree) {
  if(*tree == nil){
    return nil,tree
  }
  if((*tree).data == item){
    return *tree,tree
  }

  if(item < (*tree).data){
    return search(&(*tree).left,item)
  }
  return search(&(*tree).right,item)
}


/*
  The deletion in BST was supposed to be a bit messy, but I think
  that "return value and it's address" idea made this quite clean.

  Conventionally, here I was supposed to have a loop similar to 'insert' function.
  Rather I modified the already made functions, 'search' and 'min' to return value
  as well as address for seamless modification, as observed in line 89,95,96.
*/
func delete(tree **Tree, item int)  {
  doomed,doomed_address := search(tree,item)
  if(doomed == nil){
    fmt.Println("No such item. Can't delete!")
    return
  }

  if(doomed.left == nil && doomed.right == nil){
    *doomed_address = nil
    return
  }

  if(doomed.left != nil && doomed.right != nil){
    replacement,replacement_addr := min(&doomed.right)
    (*doomed_address).data = replacement.data
    *replacement_addr = nil
    return
  }

  if(doomed.left != nil){
    doomed.data = doomed.left.data
    doomed.left = nil
    return
  }

  if(doomed.right != nil){
    doomed.data = doomed.right.data
    doomed.right = nil
    return
  }

}

func main()  {
  root := Tree{4,nil,nil,nil}
  root_ptr := &root
  insert(&root_ptr,3,nil)
  insert(&root_ptr,9,nil)
  insert(&root_ptr,6,nil)
  insert(&root_ptr,2,nil)
  insert(&root_ptr,10,nil)
  insert(&root_ptr,12,nil)
  traverse(root_ptr)
  delete(&root_ptr,10)
  fmt.Println("Deleted one item")
  traverse(root_ptr)
}
