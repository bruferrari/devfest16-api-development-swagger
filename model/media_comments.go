package model

type MediaComments struct {
    Count int32 `json:"count,omitempty" db:"count"`
    Data array `json:"data,omitempty" db:"data"`
}
