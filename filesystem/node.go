package filesystem

type Node[T any] struct {
	IsFolder bool   `json:"isFolder"`
	Folder   Folder `json:"folder"`
	Value    *T     `json:"value"`
}
