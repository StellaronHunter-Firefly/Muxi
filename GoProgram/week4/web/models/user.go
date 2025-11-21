package models

import "time"

type User struct {
	Username  string    `json:"username" binding:"required,min=3,max=20"`
	Password  string    `json:"password" binding:"required,min=6"`
	Nickname  string    `json:"nickname" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Nickname *string `json:"nickname,omitempty"`
}
