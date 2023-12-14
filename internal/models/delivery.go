package models

type Delivery struct {
	Name    string `json:"name" validate:"required,max=30"`
	Phone   string `json:"phone" validate:"required,max=20"`
	Zip     string `json:"zip" validate:"required,max=10"`
	City    string `json:"city" validate:"required,max=50"`
	Address string `json:"address" validate:"required,max=50"`
	Region  string `json:"region" validate:"required,max=50"`
	Email   string `json:"email" validate:"required,email"`
}
