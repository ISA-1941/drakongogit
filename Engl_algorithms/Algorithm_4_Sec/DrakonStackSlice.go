// Autogenerated with DRAKON Editor 1.31

package main

import "fmt"

type Stack struct {
	data [] interface {}
}



func  (s *Stack) pop()  {
    // item 36
    if len(s.data) == 0 {
        // item 40
        fmt.Println("Stack is empty")
    } else {
        // item 39
        s.data = s.data[:len(s.data)-1]
    }
}

func  (s *Stack) push(item string)  {
    // item 30
    s.data = append(s.data, item)
}

func  main()  {
    // item 6
    stack1 := Stack{}
    stack2 := Stack{}
    // item 11
    stack1.push(10)
    stack1.push(20)
    stack1.push(30)
    stack1.push(40)
    fmt.Println("Stack -->", stack1.data)
    // item 53
    stack1.Pop()
    fmt.Println("Result for stack1--> ", stack1.data)
    stack2.Push("First")
    stack2.Push("Second")
    stack2.Push("Third")
    stack2.Push("Fourth")
    fmt.Println("Stack for stack2 -->", stack2.data)
    stack2.Pop()
    fmt.Println("Stack for stack2 -->", stack2.data)
}


