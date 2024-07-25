package main

import "fmt"

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
	t.root = insert(t.root, value)
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

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} else {
		return node // Дублирующиеся значения не допускаются
	}

	updateHeight(node)
	return balance(node)
}

func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil || node.Right == nil {
			var temp *Node
			if node.Left != nil {
				temp = node.Left
			} else {
				temp = node.Right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := minValueNode(node.Right)
			node.Value = temp.Value
			node.Right = deleteNode(node.Right, temp.Value)
		}
	}

	if node == nil {
		return node
	}

	updateHeight(node)
	return balance(node)
}

func minValueNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Value)
		inOrder(node.Right)
	}
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

func rotateRight(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func balance(node *Node) *Node {
	if node == nil {
		return node
	}

	balanceFactor := getBalance(node)

	if balanceFactor > 1 {
		if getBalance(node.Left) >= 0 {
			return rotateRight(node)
		}
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	if balanceFactor < -1 {
		if getBalance(node.Right) <= 0 {
			return rotateLeft(node)
		}
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func updateHeight(node *Node) {
	if node != nil {
		node.Height = 1 + max(height(node.Left), height(node.Right))
	}
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	values := []int{21, 17, 16, 11, 9, 7, 5, 3}
	tree := New()
	for _, value := range values {
		tree.insert(value)
	}

	fmt.Println("Pre-order traversal:")
	tree.preOrder()

	fmt.Println("In-order traversal:")
	tree.inOrder()

	fmt.Println("Post-order traversal:")
	tree.postOrder()

	fmt.Println("Deleting 16")
	tree.delete(16)

	fmt.Println("Pre-order traversal after deletion:")
	tree.preOrder()

	fmt.Println("In-order traversal after deletion:")
	tree.inOrder()

	fmt.Println("Post-order traversal after deletion:")
	tree.postOrder()
}
