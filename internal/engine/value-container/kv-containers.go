package valuecontainer

type KVContainer[T any] interface {
	Key() []byte

	Value() T
	Binary() []byte

	Compare(key []byte) bool
	Set(value T)
}
