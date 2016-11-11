package model

type Tag struct {
    MediaCount int32 `json:"media_count,omitempty" db:"media_count"`
    Name string `json:"name,omitempty" db:"name"`
}
