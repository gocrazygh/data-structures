package main

import "fmt"
//This data structure has both the properties of a stack and a queue
type List struct {
	items []int
}
//Join function appends new data
func (l *List) Join(i int) {
	l.items = append(l.items, i)
}
//RemoveFirst function removes the first value
func (l *List) RemoveFirst() int {
	toRemove := l.items[0]
	l.items = l.items[1:]
	return toRemove
}
//RemoveLast function removes the last value
func (l *List) RemoveLast() int {
	a := len(l.items)-1
	toRemove := l.items[a]
	l.items = l.items[:a]
	return toRemove
}

func main() {
	aList := List{}

	aList.Join(100)
	aList.Join(200)
	aList.Join(300)
	aList.Join(400)
	aList.Join(500)
	aList.RemoveFirst()
	aList.RemoveLast()
	fmt.Println(aList)
}