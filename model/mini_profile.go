package model

type MiniProfile struct {
    UserName string `json:"user_name,omitempty" db:"user_name"`
    FullName string `json:"full_name,omitempty" db:"full_name"`
    Id int32 `json:"id,omitempty" db:"id"`
    ProfilePicture string `json:"profile_picture,omitempty" db:"profile_picture"`
}
