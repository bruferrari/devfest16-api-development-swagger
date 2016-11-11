package model

type MediaImages struct {
    LowResolution Image `json:"low_resolution,omitempty" db:"low_resolution"`
    Thumbnail Image `json:"thumbnail,omitempty" db:"thumbnail"`
    StandardResolution Image `json:"standard_resolution,omitempty" db:"standard_resolution"`
}
