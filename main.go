package main

// LinkedList represents the collection of nodes
type LinkedList struct {
	Nodes []Node
}

// Node represents a single node in a linked list
type Node struct {
	id   int
	next int
}

func createLinkedList() LinkedList {
	return LinkedList{
		Nodes: []Node{},
	}
}

func (list LinkedList) addNode(node Node) []Node {
	newList := append(list.Nodes, node)
	return newList
}

func (list *LinkedList) deleteNode(node *Node) []Node {

}

// NewNode creates a new node
func NewNode(list *LinkedList) Node {
	return Node{
		id:   len(list.Nodes),
		next: len(list.Nodes) + 1,
	}
}

func main() {

}
