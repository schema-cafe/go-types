package filesystem

type Folder[T any] struct {
	Contents []Node[T]
}
