package models

type (
	UserProfile struct {
		ID uint64 `json:"id" binding:"required" db:"id"`
	}
)
