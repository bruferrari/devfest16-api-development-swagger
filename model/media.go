package model

type Media struct {
    CreatedTime int32 `json:"created_time,omitempty" db:"created_time"`
    Type string `json:"type,omitempty" db:"type"`
    Filter string `json:"filter,omitempty" db:"filter"`
    Tags array `json:"tags,omitempty" db:"tags"`
    Id int32 `json:"id,omitempty" db:"id"`
    User MiniProfile `json:"user,omitempty" db:"user"`
    UsersInPhoto array `json:"users_in_photo,omitempty" db:"users_in_photo"`
    Location Location `json:"location,omitempty" db:"location"`
    Comments: Media_comments `json:"comments:,omitempty" db:"comments:"`
    Likes Media_likes `json:"likes,omitempty" db:"likes"`
    Images Media_images `json:"images,omitempty" db:"images"`
    Videos Media_videos `json:"videos,omitempty" db:"videos"`
}
