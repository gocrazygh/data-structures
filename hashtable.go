package main

import (
	"fmt"
)

const aSize = 5

//hash table (array)
type HashTable struct {
	array [aSize]*Bucket
}

//bucket (linked keys)
type Bucket struct {
	head *BucketNode
}

//bucket node (stuff in the bucket)
type BucketNode struct {
	key string
	next *BucketNode
}

//insert hash table (array)
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].InsertBucket(key)
}
//search hash table (array)
func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].SearchBucket(key)
}
//delete hash table (array)
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].DeleteBucket(key)
	
}


//insert bucket (linked keys)
func (b *Bucket)InsertBucket(k string)  {
	if !b.SearchBucket(k){
	newNode := &BucketNode{key:k}
	newNode.next = b.head
	b.head = newNode
	}else {
		fmt.Println(k, "already exist")
	}
}

//search bucket (linked keys)
func (b *Bucket)SearchBucket(k string) bool {
	cNode := b.head
	for cNode != nil {
		if cNode.key == k {
			return true
		}
		cNode = cNode.next
	}
	return false
}

//delete bucket (linked keys)
func (b *Bucket)DeleteBucket(k string)  {

    if b.head.key == k{
		b.head = b.head.next
		return
	}

	pNode := b.head
	for pNode != nil{
		if pNode.next.key == k{
			pNode.next =pNode.next.next
		}
		pNode = pNode.next
	}
}

//hash
func hash(key string) int {
	r := 0
	for _, v := range key {
		r+=int(v)
	}
	return r % aSize
}

//create hash table
func Create() *HashTable {
	r := &HashTable{}
	for i := range r.array {
		r.array[i] = &Bucket{}
	}
	return r
}

func main() {
	myHashTable := Create()

	keys := []string{
		"mike",
		"sam",
		"ben",
		"ham",
		"jake",
	}

	for _, v := range keys {
		myHashTable.Insert(v)
	}
    
	fmt.Println(myHashTable.Search("mike"))
	myHashTable.Delete("mike")
	fmt.Println(myHashTable.Search("mike"))
}