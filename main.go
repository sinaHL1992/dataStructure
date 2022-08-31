package main

import "fmt"

const hashSize = 7

type hashTable struct {
	array [hashSize]*bucket
}
type Node struct {
	key      string
	nextNode *Node
}
type bucket struct {
	headNode *Node
}

func (hashTable *hashTable) insert(name string) {
	index := hash(name)
	hashTable.array[index].Insert(name)

}

func (hashTable *hashTable) delete(name string) {
	index := hash(name)
	hashTable.array[index].Delete(name)

}
func (hashTable hashTable) search(name string) {
	index := hash(name)
	hashTable.array[index].Search(name)
}

func (b bucket) Search(name string) bool {
	var node *Node
	for node = b.headNode; node != nil; node = node.nextNode {
		if node.key == name {
			return true
		}
	}
	return false
}
func (bucket *bucket) Insert(name string) {
	if !bucket.Search(name) {
		node := &Node{key: name}
		node.nextNode = bucket.headNode
		bucket.headNode = node
	} else {
		fmt.Println("Already Exists")
	}

}
func (b *bucket) Delete(name string) {
	if b.headNode.key == name {
		b.headNode = b.headNode.nextNode
	}
	var node *Node
	for node = b.headNode; node != nil; node = node.nextNode {
		if node.nextNode.key == name {
			node.nextNode = node.nextNode.nextNode
		}
		node = node.nextNode
	}

}

func hash(name string) int {
	sum := 0
	for _, v := range name {
		sum = sum + int(v)
	}
	return sum % hashSize
}

func Init() *hashTable {
	result := hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return &result
}
func main() {
	myhash := Init()
	fmt.Println(myhash)
	myhash.insert("EHSAN")
	myhash.insert("AMELIE")
	myhash.insert("ELAHE")
	myhash.insert("RAMIN")
	myhash.insert("SINA")
	myhash.insert("SHIVA")

}
