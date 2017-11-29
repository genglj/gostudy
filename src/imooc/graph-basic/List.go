// List based on slice
package graph_basic

import (
	"strings"
	"fmt"
)
const (
	initCapacity = 16
	growthFactor = float32(2.0)  // growth by 100%
)

// List holds the elements in a slice
type List struct {
	elements []interface{}
	size     int
}

//New instantiates a new empty list
func NewList() *List {
	return &List{}
}

//Size returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

//Empty returns true if list doesn't contain any elements.
func (list *List) Empty() bool {
	return list.size == 0
}

//Clear removes all elements from the list
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

//Add appends a value at the end of the list
func (list *List) Add(value interface{}) {
	list.resize()
	list.elements[list.size] = value
	list.size++
}

//Get returns the element at index
//Second return parameter is true if index is within bounds of the array and the array is not empty,otherwise false
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Check that the index is within bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) String() string {
	str := "List{"
	values := []string{}
	for _,value := range list.elements[0:list.size] {
		values = append(values,fmt.Sprintf("%v",value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) resize() {
	currentCapacity := cap(list.elements)
	if list.size + 1 >  currentCapacity {
		if currentCapacity == 0 {
			currentCapacity = initCapacity/int(growthFactor)
		}
		newElements := make([]interface{}, currentCapacity*int(growthFactor), currentCapacity*int(growthFactor))
		copy(newElements, list.elements)
		list.elements = newElements
	}
}

