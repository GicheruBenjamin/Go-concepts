package object

import "fmt"

// Object represents a generic object in memory
type Object struct {
	ID   int
	Data []byte // Simulates memory-intensive data
}

// ObjectManager manages the lifecycle of objects
type ObjectManager struct {
	counter int
	pool    []*Object
}

// NewObjectManager creates a new ObjectManager
func NewObjectManager() *ObjectManager {
	return &ObjectManager{
		pool: make([]*Object, 0),
	}
}

// CreateObject creates a new object with simulated memory usage
func (om *ObjectManager) CreateObject(size int) *Object {
	om.counter++
	obj := &Object{
		ID:   om.counter,
		Data: make([]byte, size), // Allocate memory
	}
	om.pool = append(om.pool, obj)
	fmt.Printf("Created Object ID: %d, Memory Allocated: %d bytes\n", obj.ID, size)
	return obj
}

// DestroyObject removes an object from the pool and dereferences it
func (om *ObjectManager) DestroyObject(obj *Object) {
	for i, o := range om.pool {
		if o.ID == obj.ID {
			om.pool = append(om.pool[:i], om.pool[i+1:]...)
			break
		}
	}
	// Explicitly set to nil to allow garbage collection
	fmt.Printf("Destroyed Object ID: %d\n", obj.ID)
	obj = nil
}
