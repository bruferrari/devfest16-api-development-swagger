package model

type InlineResponse2007 struct {
    Meta inline_response_200_7_meta `json:"meta,omitempty" db:"meta"`
    Data array `json:"data,omitempty" db:"data"`
}
