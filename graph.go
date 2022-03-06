//Here's a simple graph data structure
package main

import "fmt"

//Here is the structure for the graph containg the vertex
type Graph struct {
	vertices []*Vertex
}

//The vertex are the key data point
type Vertex struct {
	key int
	ajecent []*Vertex
}

//This function adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	//The error avoids duplicate vertex
	if contains(g.vertices, k) {
		err := fmt.Errorf("cannot add because vertex %v's already exist", k)
		fmt.Println(err.Error())
	}else {
	//This appends an vertex
	g.vertices = append(g.vertices, &Vertex{key: k})	
	}
}
//This function adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	//Variables for where the edges are
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		//Tis error avoids creating edges that don't exist
		err := fmt.Errorf("(%v ---> %v) is an invalid edge", from, to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.ajecent, to){
				//Tis error avoids creating edges that already exist
		err := fmt.Errorf("(%v ---> %v) edge already exist", from, to)
		fmt.Println(err.Error())
	}else {
		//Add edge
		fromVertex.ajecent = append(fromVertex.ajecent, toVertex)
	}
}

//Get the vertex
func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key ==k {
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*Vertex,k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//This print's the vertex
func (g *Graph) Print()  {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.key)
		for _, v := range v.ajecent {
			fmt.Printf(" %v ", v.key)			
		}
	}
	fmt.Println()
}

func main() {
	myGraph := &Graph{}//Creating a new graph
	for i := 0; i < 10; i++ {//Creating 10 vertices
		myGraph.AddVertex(i)
	}

	myGraph.AddVertex(10)//Add a vertex
	myGraph.AddEdge(5,7)//Add an edge
	myGraph.AddEdge(15,5)//Add an invalid edge
	myGraph.AddEdge(7,5)//Add a edge
	myGraph.AddEdge(7,5)//Add a duplicate edge 

	myGraph.Print()
}
