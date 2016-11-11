package model

type Image struct {
    Width int32 `json:"width,omitempty" db:"width"`
    Height int32 `json:"height,omitempty" db:"height"`
    Url string `json:"url,omitempty" db:"url"`
}
