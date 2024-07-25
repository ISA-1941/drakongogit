// Autogenerated with DRAKON Editor 1.31

package main
import (
"fmt"
)

var Fnode bool = false
 
type BSTree struct {
	rootValue int
	leftNode  *BSTree
	rightNode *BSTree
}







func  (root *BSTree) InOrder()  {
    // item 98
    if root == nil {
        // item 101
        return
    } else {
        
    }
    // item 102
    root.leftNode.InOrder()
    fmt.Printf("%d ", root.rootValue)
    root.rightNode.InOrder()
}

func  (root *BSTree) Insert(value int)  {
    // item 42
    if value < root.rootValue {
        // item 45
        if root.leftNode == nil {
            // item 48
            root.leftNode =
             createNode(value)
            // item 49
            return
        } else {
            
        }
        // item 50
        root.leftNode.Insert(value)
    } else {
        
    }
    // item 51
    if value > root.rootValue {
        // item 54
        if nil == root.rightNode {
            // item 57
            root.rightNode =
             createNode(value)
            // item 58
            return
        } else {
            
        }
        // item 59
        root.rightNode.Insert(value)
    } else {
        
    }
}

func  (root *BSTree) LevelWiseTraversal()  {
    // item 221
    nodeList := make([]*BSTree, 0)
    // item 222
    if nil == root {
        // item 225
        return
    } else {
        
    }
    // item 226
    temp := &BSTree{}
    nodeList = append(nodeList, root)
    for {
        // item 227
        if len(nodeList) > 0 {
            
        } else {
            break
        }
        // item 229
        temp, nodeList = 
        nodeList[0], nodeList[1:]
        // item 2300001
        if temp.leftNode != nil {
            // item 236
            nodeList = 
            append(nodeList, 
            temp.leftNode)
        } else {
            // item 2300002
            if temp.rightNode != nil {
                // item 237
                nodeList = 
                append(nodeList, 
                temp.rightNode)
            } else {
                
            }
        }
    }
    // item 241
    fmt.Println()
}

func  (root *BSTree) PostOrder()  {
    // item 64
    if nil == root {
        // item 67
        return
    } else {
        
    }
    // item 68
    fmt.Printf("%d ", root.rootValue)
    root.leftNode.InOrder()
    root.rightNode.InOrder()
}

func  (root *BSTree) PreOrder()  {
    // item 343
    if root == nil {
        // item 346
        return
    } else {
        
    }
    // item 347
    root.leftNode.InOrder()
    root.rightNode.InOrder()
    fmt.Printf("%d ", root.rootValue)
}

func  createNode(value int) *BSTree {
    // item 11
    return &BSTree{
    		rootValue: value,
    		leftNode:  nil,
    		rightNode: nil,
    	}
}

func  deleteNode(root *BSTree, val int) *BSTree {
    // item 1620001
    if root.rootValue > val {
        // item 170
        root.leftNode =
         deleteNode(root.leftNode, val)
    } else {
        // item 1620002
        if root.rootValue < val {
            // item 171
            root.rightNode =
            deleteNode(root.
            rightNode, val)
        } else {
            // item 1620003
            if root.rootValue == val {
                
            } else {
                // item 1620004
                panic("Not expected condition.")
            }
            // item 180
            L1 := root.leftNode == nil
            R1 := root.rightNode == nil
            L2 := root.leftNode != nil
            R2 := root.rightNode != nil
            // item 1810001
            if L1 && R1 {
                // item 188
                root = nil
                // item 189
                return root
                // item 199
                tempNode := minValued(root.rightNode)
                root.rootValue = tempNode.rootValue
                root.rightNode =
                // item 351
                deleteNode(root.rightNode, 
                tempNode.rootValue)
            } else {
                // item 1810002
                if L1 && R2 {
                    // item 190
                    temp := root.rightNode
                    // item 191
                    root = nil
                    // item 192
                    root = temp
                    // item 193
                    return root
                } else {
                    // item 1810003
                    if L2 && R1 {
                        // item 197
                        temp := root.leftNode
                        // item 194
                        root = nil
                        // item 195
                        root = temp
                        // item 196
                        return root
                    } else {
                        // item 1810004
                        if L2 && R2 {
                            
                        } else {
                            // item 1810005
                            panic("Not expected condition.")
                        }
                    }
                }
            }
        }
    }
    // item 200
    return root
}

func  findNode(root *BSTree, val int) *BSTree {
    // item 348
    fmt.Println("findNode -->", root)
    fmt.Println("rootValue -->", root.rootValue)
    // item 2940001
    if root.rootValue > val {
        // item 302
        root.leftNode =
         findNode(root.leftNode, val)
    } else {
        // item 2940002
        if root.rootValue < val {
            // item 303
            root.rightNode =
             findNode(root.rightNode, val)
        } else {
            // item 2940003
            if root.rootValue == val {
                
            } else {
                // item 2940004
                panic("Not expected condition.")
            }
            // item 332
            Fnode = true
        }
    }
    // item 325
    return root
}

func  main()  {
    // item 273
    numbers := []int{56, 38, 73, 
     25, 64, 15, 87, 47, 93}
    var root *BSTree
    var rootNode bool = true
    for _, value := range numbers {
        // item 276
        if rootNode {
            // item 279
            root = createNode(value)
            rootNode = false
        } else {
            
        }
        // item 280
        root.Insert(value)
    }
    // item 282
       root.PreOrder()
    fmt.Println()
    root.InOrder()
       fmt.Println()
       root.PostOrder()
    fmt.Println()
    // item 334
    val := 93
    findNode(root, val)
    // item 335
    if Fnode == true {
        // item 338
        fmt.Println("\n The node ", val,
        "is exist")
    } else {
        // item 339
        fmt.Println("\n The node  ", val,
        "does not exist")
    }
    // item 285
    minValued(root)
    maxValued(root)
    // item 286
    root.LevelWiseTraversal()
    // item 287
    deleteNode(root, val)
    // item 288
    root.PostOrder()
}

func  maxValued(root *BSTree) *BSTree {
    // item 108
    temp := root
    for {
        // item 129
        if nil != temp && 
temp.rightNode != nil {
            
        } else {
            break
        }
        // item 131
        temp = temp.rightNode
    }
    // item 111
    fmt.Println("Max temp = ", temp)
    // item 112
    return temp
}

func  minValued(root *BSTree) *BSTree {
    // item 75
    temp := root
    for {
        // item 76
        if nil != temp && 
temp.leftNode != nil {
            
        } else {
            break
        }
        // item 126
        temp = temp.leftNode
    }
    // item 80
    fmt.Println("Min temp = ", temp)
    // item 81
    return temp
}


