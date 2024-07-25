package main
import "fmt"
type Node struct {
    value int
    next *Node
}
type Graph struct {
    vertices []*Node
}
func (g *Graph) AddVertex(value int) {
    newNode := &Node{value: value}
    g.vertices = append(g.vertices, newNode)
}

func (g *Graph) AddEdge(v1, v2 int) {
    node1 := g.vertices[v1]
    node2 := g.vertices[v2]

    newNode := &Node{value: v2}
    newNode.next = node1.next
    node1.next = newNode

    newNode = &Node{value: v1}
    newNode.next = node2.next
    node2.next = newNode
}

func (g *Graph) GetNeighbors(vertex int) {
    node := g.vertices[vertex]
    for node != nil {
        fmt.Println(node.value)
        node = node.next
    }
}

func main() {
    graph := Graph{}
    graph.AddVertex(0)
    graph.AddVertex(1)
    graph.AddVertex(2)

    graph.AddEdge(0, 1)
    graph.AddEdge(1, 2)

    graph.GetNeighbors(0)
}
