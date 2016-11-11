package model

type Location struct {
    Id string `json:"id,omitempty" db:"id"`
    Name string `json:"name,omitempty" db:"name"`
    Latitude float32 `json:"latitude,omitempty" db:"latitude"`
    Longitude float32 `json:"longitude,omitempty" db:"longitude"`
}
