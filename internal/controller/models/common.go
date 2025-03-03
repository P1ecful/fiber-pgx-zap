package models

type GetManyResult[T any] struct {
	Count  uint64 `json:"count"`
	Result []T    `json:"result"`
}
