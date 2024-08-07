// Autogenerated with DRAKON Editor 1.31

package main

import "fmt"

type Node struct {
	data string
	nextNode *Node
}
type LinkedList struct {
	len      int
	headNode *Node // тип Указатель адреса  структуры узла Uzel
	n int
}



func  (ll *LinkedList) iterateList() *Node {
    // item 97
    var node *Node
    // item 980001
    node = ll.headNode;
    for {
        // item 980002
        if node != nil {
            
        } else {
            break
        }
        // item 100
        fmt.Println(node.data)
        // item 980003
        node = node.nextNode;
    }
    // item 103
    return node
}

func  (ll *LinkedList) pushBack(val string)  {
    // item 72
    node := ll.headNode
    newNode := &Node{data: val, 
    nextNode: nil}
    // item 73
    if node == nil {
        // item 77
        return
    } else {
        
    }
    for {
        // item 78
        if node.nextNode != nil {
            
        } else {
            break
        }
        // item 80
        node = node.nextNode
    }
    // item 76
    node.nextNode = newNode
}

func  (ll *LinkedList) pushFront(val string)  {
    // item 87
    var node = &Node{data: val, nextNode: nil}
    // item 88
    if ll.headNode != nil {
        // item 91
        node.nextNode = ll.headNode
    } else {
        
    }
    // item 92
    ll.headNode = node
    // item 93
    ll.len++
}

func  (ll *LinkedList) removeVal(val string) *Node {
    // item 32
    node := ll.headNode
    for {
        // item 40
        if node != nil && node.data == val {
            
        } else {
            break
        }
        // item 33
        ll.headNode =
         node.nextNode
        // item 43
        node = ll.headNode
    }
    for {
        // item 44
        if node != nil {
            
        } else {
            break
        }
        // item 34
        nextNode := node.nextNode
        // item 35
        if nextNode != nil && 
nextNode.data == val {
            // item 36
            node.nextNode =
            nextNode.nextNode
        } else {
            // item 37
            node = nextNode
        }
    }
    // item 38
    return node
}

func  (ll *LinkedList) searchData(val string)  {
    // item 54
    var node *Node
    // item 550001
    node = ll.headNode;
    for {
        // item 550002
        if node != nil {
            
        } else {
            break
        }
        // item 57
        if node.data ==  val {
            // item 60
            ll.n++
        } else {
            
        }
        // item 550003
        node = node.nextNode;
    }
    // item 61
    if ll.n != 0 {
        // item 65
        fmt.Println("The desired value ", val, 
        "occurs ", ll.n, "times")
    } else {
        // item 64
        fmt.Println("The desired value ", val, 
        "does not exist")
    }
    // item 66
    return
}

func  main()  {
    // item 6
    var ll LinkedList = LinkedList{}
    ll.n = 0
    ll.pushFront("Smith J.")
    ll.pushBack("Brown G.")
    ll.pushBack("Shafler P.")
    ll.pushBack("Wiley S.")
    ll.pushBack("Wiley S.")
    ll.pushBack("Atallah N.")
    ll.searchData("Wiley S.")
    ll.removeVal("Wiley S.")
    ll.iterateList()
    fmt.Println(" Wiley S. is deleted")
}


