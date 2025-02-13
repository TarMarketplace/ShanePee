package dto

type ArrayResponse[T any] struct {
	Data []*T `json:"data"`
}
