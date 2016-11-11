package model

type InlineResponse200 struct {
    Data array `json:"data,omitempty" db:"data"`
}
