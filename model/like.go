package model

type Like struct {
    UserName string `json:"user_name,omitempty" db:"user_name"`
    FirstName string `json:"first_name,omitempty" db:"first_name"`
    LastName string `json:"last_name,omitempty" db:"last_name"`
    Type string `json:"type,omitempty" db:"type"`
    Id string `json:"id,omitempty" db:"id"`
}
