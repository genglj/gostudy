package graph_basic

type Component struct {
	graph   *SparseGraph
	visited []bool
	count   int
}

func NewComponent(graph *SparseGraph) *Component {
	visited := make([]bool, graph.vertexNum)
	return &Component{
		graph:   graph,
		visited: visited,
		count:   0,
	}
}

func (component *Component) Count() int {
	for i := 0; i < component.graph.vertexNum; i++ {
		if !component.visited[i] {
			component.count++
			component.deepFirstSearch(i)
		}
	}
	return component.count
}

func (component *Component) deepFirstSearch(v int) {
	component.visited[v] = true
	component.graph.Do(func(value int) {
		if !component.visited[value] {
			component.deepFirstSearch(value)
		}
	}, v)

}
