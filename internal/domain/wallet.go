package domain

import "time"

type Wallet struct {
    ID          string    `json:"id" bson:"_id,omitempty"`
    UserID      string    `json:"user_id" bson:"user_id"`
    Amount      float64   `json:"amount" bson:"amount"`
    Type        string    `json:"type" bson:"type"` // income/expense
    Category    string    `json:"category" bson:"category"`
    Description string    `json:"description" bson:"description"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}