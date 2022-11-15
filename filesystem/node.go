package filesystem

type Node[T any] struct {
	IsFolder bool
	Folder   Folder[T]
	Value    T
}
