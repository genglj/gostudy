package graph_basic

import (
"fmt"
"testing"
)

func TestComponent(t *testing.T) {
	sg := NewSparseGraph(10, false)
	//fmt.Println(sg.EdgeNum(), sg.VertexNum())
	sg.AddEdge(2, 4)
	sg.AddEdge(6, 9)
	sg.AddEdge(6, 8)
	//fmt.Println(dp.EdgeNum())
	cp := NewComponent(sg)

	for i:=0;i<cp.graph.vertexNum;i++{
		fmt.Printf("%d: ",i)
		cp.graph.Do(func(value int){
			fmt.Printf("%d ",value)
		},i)
		fmt.Println()
	}

	fmt.Println(cp.Count())
}