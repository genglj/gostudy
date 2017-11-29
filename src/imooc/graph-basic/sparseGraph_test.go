package graph_basic

import (
	"fmt"
	"testing"
)

func TestSparseGraph(t *testing.T) {
	sg := NewSparseGraph(10, false)
	//fmt.Println(sg.EdgeNum(), sg.VertexNum())
	sg.AddEdge(2, 4)
	sg.AddEdge(6, 9)
	sg.AddEdge(6, 8)
	//fmt.Println(dp.EdgeNum())
	sg.Do(func(v int) {
		fmt.Printf("%d ",v)
	},6)
}