package model

type InlineResponse2005 struct {
    Meta inline_response_200_4_meta `json:"meta,omitempty" db:"meta"`
    Data object `json:"data,omitempty" db:"data"`
}
