package lfqueue

type Queue interface {
	Enqueue(v interface{}) error
	Dequeue() interface{}
	Length() uint32
}
