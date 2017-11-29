// dense graph using adjacency matrix
package graph_basic

type DenseGraph struct {
	//Vertex number
	vertexNum int
	//Edge number
	edgeNum int
	//Whether this graph is a directed graph
	directed bool
	//Matrix
	g [][]bool
}

// n Vertex number
//directed Whether this graph is a directed graph
func NewDenseGraph(n int, directed bool) *DenseGraph {
	list := make([][]bool, n)
	for i := 0; i < n; i++ {
		list[i] = make([]bool, n)
	}
	return &DenseGraph{
		vertexNum:        n,
		edgeNum:        0,
		directed: directed,
		g:        list,
	}
}
// return this graph's vertex number
func (dg DenseGraph) VertexNum() int {
	return dg.vertexNum
}
//return this graph's edge number
func (dg DenseGraph) EdgeNum() int {
	return dg.edgeNum
}
//Add an edge from vertex p to vertex q
func (dg *DenseGraph) AddEdge(p,q int) {
	if dg.HasEdge(p,q) {
		return
	}
	dg.g[p][q] = true
	if !dg.directed {
		dg.g[q][p] = true
	}
	dg.edgeNum++
}

//HasEdge returns true if exists an edge between vertex p and vertex q,otherwise false
func (dg DenseGraph) HasEdge(p,q int) bool{
	return dg.g[p][q]
}



