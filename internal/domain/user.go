package domain

import "time"

type User struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Username  string    `json:"username" bson:"username"`
    Password  string    `bson:"password"`
    Email     string    `json:"email" bson:"email"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}