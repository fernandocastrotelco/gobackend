package main

import "fmt"

// ArraySize es el tamaño q le vamos a dar a mi hashmap, por defecto 7
const ArraySize = 7

// HashTable de array por ejemplo
type HashTable struct {
	array []*bucket
}

// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	fmt.Printf("index: %d", index)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	fmt.Printf("index: %d", index)
	return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	fmt.Printf("index: %d", index)
	h.array[index].delete(key)
}

// bucket es una linkedlist con la info a guardar
// se implementa una linkedlist para manejar las colisiones
type bucket struct {
	head *bucketNode
}

// bucketNode es el cuerpo de la linkedList
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert
func (b *bucket) insert(k string) {
	if b.search(k) {
		fmt.Printf("%s ya existe!", k)
		return
	}
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search busca una key en la linkedList
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % 3
}

// NewHashMap va a crear una bucket en cada posicion del array q contiene nuestro hashmap
func NewHashMap(s int) *HashTable {
	arrs := make([]*bucket, s)
	result := &HashTable{array: arrs}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := NewHashMap(3)
	fmt.Print(testHashTable)

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	fmt.Println("insertamos los nombres")

	for _, v := range list {
		testHashTable.Insert(v)
	}

	fmt.Println("posicion 0 ", testHashTable.array[0].head.key, testHashTable.array[0].head.next.key)
	fmt.Println("posicion 1 ", testHashTable.array[1].head.key, testHashTable.array[1].head.next.key)
	fmt.Println("posicion 2 ", testHashTable.array[2].head.key)

	fmt.Println("buscamos a RANDY: ", testHashTable.Search("RANDY"))

	fmt.Println("eliminamos a RANDY: ")
	testHashTable.Delete("RANDY")

	fmt.Println("buscamos a RANDY: ", testHashTable.Search("RANDY"))

	fmt.Println("posicion 0 ", testHashTable.array[0].head.key, testHashTable.array[0].head.next.key)
	fmt.Println("posicion 1 ", testHashTable.array[1].head.key, testHashTable.array[1].head.next.key)
	fmt.Println("posicion 2 ", testHashTable.array[2].head.key)
}
