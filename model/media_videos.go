package model

type MediaVideos struct {
    LowResolution Image `json:"low_resolution,omitempty" db:"low_resolution"`
    StandardResolution Image `json:"standard_resolution,omitempty" db:"standard_resolution"`
}
