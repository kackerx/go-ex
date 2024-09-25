package set

type Set interface {
	Add(e int)
	Remove(e int)
	Contains(e int)
	Size() int
	IsEmpty() bool
}
