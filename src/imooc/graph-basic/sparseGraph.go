//sparse graph using adjacency list
package graph_basic

type SparseGraph struct {
	//Vertex number
	vertexNum int
	//Edge number
	edgeNum int
	//Whether this graph is a directed graph
	directed bool
	//List
	g []List
}

// n Vertex number
//directed Whether this graph is a directed graph
func NewSparseGraph(n int, directed bool) *SparseGraph {
	g := make([]List, n)
	for i := 0; i < n; i++ {
		g[i] = *NewList()
	}
	return &SparseGraph{
		vertexNum: n,
		edgeNum:   0,
		directed:  directed,
		g:         g,
	}
}

// return this graph's vertex number
func (sg SparseGraph) VertexNum() int {
	return sg.vertexNum
}

//return this graph's edge number
func (sg SparseGraph) EdgeNum() int {
	return sg.edgeNum
}

//Add an edge from vertex p to vertex q
func (sg *SparseGraph) AddEdge(p, q int) {
	sg.g[p].Add(q)
	if p != q && !sg.directed {
		sg.g[q].Add(p)
	}
	sg.edgeNum++
}

//HasEdge returns true if exists an edge between vertex p and vertex q,otherwise false
func (dg SparseGraph) HasEdge(p, q int) bool {
	return false
}

//The sparse graph iterator
func (sg SparseGraph) Do(fn func(int), e int) {
	for i:=0;i<sg.g[e].Size();i++ {
		value,_ := sg.g[e].Get(i)
		fn(value.(int))
	}
}
