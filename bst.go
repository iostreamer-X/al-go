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
  if(tree == nil){
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
