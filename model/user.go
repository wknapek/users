package model

type User struct {
	ID         string `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required,min=2"`
	SignUpTime int64  `json:"signupTime" validate:"required,min=-3786825600000"`
}
