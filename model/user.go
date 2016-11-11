package model

type User struct {
    Id int32 `json:"id,omitempty" db:"id"`
    Username string `json:"username,omitempty" db:"username"`
    FullName string `json:"full_name,omitempty" db:"full_name"`
    ProfilePicture string `json:"profile_picture,omitempty" db:"profile_picture"`
    Bio string `json:"bio,omitempty" db:"bio"`
    Website string `json:"website,omitempty" db:"website"`
    Counts User_counts `json:"counts,omitempty" db:"counts"`
}
