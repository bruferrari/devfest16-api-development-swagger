package model

type UserCounts struct {
    Media int32 `json:"media,omitempty" db:"media"`
    Follows int32 `json:"follows,omitempty" db:"follows"`
    FollwedBy int32 `json:"follwed_by,omitempty" db:"follwed_by"`
}
