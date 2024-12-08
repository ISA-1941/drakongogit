// Autogenerated with DRAKON Editor 1.31

package main

import (  
   "fmt"    
)


var fBool bool = false

var temp, tmp *Node

type Node struct {
    Value, Height int
    Left, Right *Node
}

type Tree struct {
    root *Node
}
func New() *Tree {
 return &Tree{}
}

func (t *Tree) insert(value int) {
    t.root = insertNode(t.root, value)
}

func  (tree *Tree) find(value int) *Node {
    // item 250
    return findNode(tree.root, value)
}

func (t *Tree) delete(value int) {
    t.root = deleteNode(t.root, value)
}

func (t *Tree) preOrder() {
    preOrder(t.root)
    fmt.Println()
}

func (t *Tree) inOrder() {
    inOrder(t.root)
    fmt.Println()
}

func (t *Tree) postOrder() {
    postOrder(t.root)
    fmt.Println()
}




func  balance(node *Node) *Node {
    // item 403
    updateHeight(node)
    // item 404
    bF:= gB(node)
    // item 4120001
    if bF > 1 &&
 gB(node.Left) >= 0 {
        // item 405
        return rotateRight(node)
    } else {
        // item 4120002
        if bF < -1 &&
gB(node.Left) < 0 {
            // item 406
            return rotateLeft(node)
        } else {
            // item 4120003
            if bF > 1 && 
gB(node.Right) <= 0 {
                // item 408
                node.Left = 
                rotateLeft(node.Left)
                // item 407
                return rotateRight(node)
            } else {
                // item 4120004
                if bF < -1 && 
gB(node.Right) > 0 {
                    // item 409
                    node.Right = 
                    rotateRight(node.Right)
                    // item 410
                    return rotateLeft(node)
                } else {
                    
                }
            }
        }
    }
    // item 411
    return node
}

func  deleteNode(node *Node, value int) *Node {
    // item 336
    if node == nil {
        // item 339
        return node
    } else {
        
    }
    // item 3650001
    if value < node.Value {
        // item 340
        node.Left = 
        deleteNode(node.Left, value)
    } else {
        // item 3650002
        if value > node.Value {
            // item 341
            node.Right = 
            deleteNode(node.Right, value)
        } else {
            // item 3650003
            if node.Left == nil || 
node.Right == nil {
                // item 343
                if node.Left != nil {
                    // item 346
                     temp = node.Left
                } else {
                    // item 347
                    temp = node.Right
                }
                // item 348
                if temp == nil {
                    // item 349
                    temp = node
                    // item 350
                    node = nil
                } else {
                    // item 351
                    *node = *temp
                }
            } else {
                // item 3650004
                if node.Left != nil && 
node.Right != nil {
                    
                } else {
                    // item 3650005
                    panic("Not expected condition.")
                }
                // item 352
                tmp = minValueNode(node.Right)
                // item 353
                node.Value = tmp.Value
                // item 354
                node.Right = 
                deleteNode(node.Right, tmp.Value)
            }
        }
    }
    // item 378
    if node == nil {
        // item 381
        return node
    } else {
        
    }
    // item 382
    updateHeight(node)
    // item 383
    return balance(node)
}

func  findNode(node *Node, value int) *Node {
    // item 2270001
    if node == nil {
        // item 220
        return nil
    } else {
        // item 2270002
        if node.Value == value {
            // item 222
            fBool = true
            // item 221
            return node
        } else {
            // item 2270003
            if value > node.Value {
                // item 223
                return findNode(node.Right, value)
            } else {
                // item 2270004
                if value < node.Value {
                    
                } else {
                    // item 2270005
                    panic("Not expected condition.")
                }
                // item 224
                return findNode(node.Left, value)
            }
        }
    }
    // item 225
    return nil
}

func  gB(node *Node) int {
    // item 127
    if node == nil {
        // item 130
        return 0
    } else {
        
    }
    // item 131
    return height(node.Left) -
     height(node.Right)
}

func  height(node *Node) int {
    // item 199
    if node == nil {
        // item 202
        return 0
    } else {
        
    }
    // item 203
    return node.Height
}

func  inOrder(node *Node)  {
    // item 66
    if node != nil {
        // item 81
        inOrder(node.Left)
        // item 69
        fmt.Printf("%d ", node.Value)
        // item 70
        inOrder(node.Right)
    } else {
        
    }
}

func  insertNode(node *Node, value int) *Node {
    // item 87
    if node == nil {
        // item 90
        return &Node{Value: value, 
        Height: 1 }
    } else {
        
    }
    // item 910001
    if value < node.Value {
        // item 99
        node.Left = 
        insertNode(node.Left, value)
    } else {
        // item 910002
        if value > node.Value {
            // item 100
            node.Right = 
            insertNode(node.Right, value)
        } else {
            // item 98
            return node
        }
    }
    // item 102
    updateHeight(node)
    // item 103
    return balance(node)
}

func  main()  {
    // item 32
    values := [] int {21,17,16,
    11,9,7,5,3}
    // item 33
    tree := New()
    for _, value := range values {
        // item 36
        tree.insert(value)
    }
    // item 38
    fmt.Println("Pre-order traversal:")
    tree.preOrder()
    // item 39
    fmt.Println("In-order traversal:")
     tree.inOrder()
    // item 40
    fmt.Println("Post-order traversal:")
     tree.postOrder()
    // item 236
    numF := 3
    tree.find(numF)
    // item 239
    if fBool == true {
        // item 237
        fmt.Println("\n The number: ", 
        numF, "is present")
    } else {
        // item 238
        fmt.Println("\n The number: ", 
        numF, "is missing")
    }
    // item 392
    numR := 16
    fmt.Println("Removing number:", numR )
    tree.delete(numR)
    // item 391
    tree.inOrder()
    // item 396
    fmt.Println("The number:", numR, "removed")
}

func  max(a,b int) int {
    // item 46
    if a > b {
        // item 49
        return a
    } else {
        // item 50
        return b
    }
}

func  minValueNode(node *Node) *Node {
    // item 265
    current := node
    for {
        // item 266
        if current.Left != nil {
            
        } else {
            break
        }
        // item 269
        current = current.Left
    }
    // item 271
    return current
}

func  postOrder(node *Node)  {
    // item 76
    if node != nil {
        // item 80
        postOrder(node.Left)
        postOrder(node.Right)
        // item 79
        fmt.Printf("%d ", node.Value)
    } else {
        
    }
}

func  preOrder(node *Node)  {
    // item 56
    if node != nil {
        // item 59
        fmt.Printf("%d ", node.Value)
        // item 60
        preOrder(node.Left)
        preOrder(node.Right)
    } else {
        
    }
}

func  rotateLeft(x *Node) *Node {
    // item 118
    y := x.Right
    t := y.Left
    // item 119
    y.Left = x
    x.Right = t
    // item 120
    updateHeight(x)
    updateHeight(y)
    // item 121
    return y
}

func  rotateRight(y *Node) *Node {
    // item 109
    x := y.Left
    t := x.Right
    // item 110
    x.Right = y
    y.Left = t
    // item 111
    updateHeight(y)
    updateHeight(x)
    // item 112
    return x
}

func  updateHeight(node *Node)  {
    // item 256
    if node != nil {
        // item 259
        node.Height = 1 + 
        max(height(node.Left), 
        height(node.Right))
    } else {
        
    }
}

