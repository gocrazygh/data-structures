//The hash table is famous for it's constant time complexity

package main

import (
	"fmt"
)

const aSize = 5

//The structure of the hash table is an array
type HashTable struct {
	array [aSize]*Bucket
}

//The bucket is the linked list appeded to the hash table
type Bucket struct {
	head *BucketNode
}

//The keys will be stored in the bucket and will be a bucket node
type BucketNode struct {
	key string
	next *BucketNode
}

//The insert function for the hash table (array)
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].InsertBucket(key)
}
//The search function for the hash table (array)
func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].SearchBucket(key)
}
//The delete function for the hash table (array)
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].DeleteBucket(key)
	
}

//These three functions edit the bucket
//The insert function for the bucket(linked keys)
func (b *Bucket)InsertBucket(k string)  {
	if !b.SearchBucket(k){
	newNode := &BucketNode{key:k}
	newNode.next = b.head
	b.head = newNode
	}else {
		fmt.Println(k, "already exist")
	}
}

//The search function for the bucket (linked keys)
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

//The delete function for the bucket (linked keys)
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

//This function turns the string to ascii int and divided by the size
//of the array to get it's index
func hash(key string) int {
	r := 0
	for _, v := range key {
		r+=int(v)
	}
	return r % aSize
}

//This function creates a hash table
func Create() *HashTable {
	r := &HashTable{}
	for i := range r.array {
		r.array[i] = &Bucket{}
	}
	return r
}

func main() {
	myHashTable := Create()//Creating a hash table

	keys := []string{
		"mike",
		"sam",
		"ben",
		"ham",
		"jake",
	}

	for _, v := range keys {//Inserting keys to the hash table
		myHashTable.Insert(v)
	}
    
	fmt.Println(myHashTable.Search("mike"))//Searching for mike
	myHashTable.Delete("mike")//Deleting mike
	fmt.Println(myHashTable.Search("mike"))//Searching again for mike
}
