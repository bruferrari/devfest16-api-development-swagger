package model

type InlineResponse2004 struct {
    Meta inline_response_200_4_meta `json:"meta,omitempty" db:"meta"`
    Data array `json:"data,omitempty" db:"data"`
}
