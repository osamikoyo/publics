package entity

type UID uint64

type Element struct {
	ID      UID
	Parents []Topic
	Self    *Topic
}
