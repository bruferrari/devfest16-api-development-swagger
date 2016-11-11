package model

type Comment struct {
    Id string `json:"id,omitempty" db:"id"`
    CreatedTime string `json:"created_time,omitempty" db:"created_time"`
    Text string `json:"text,omitempty" db:"text"`
    From MiniProfile `json:"from,omitempty" db:"from"`
}
