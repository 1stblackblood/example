package entity

import(
	"gorm.io/gorm"
)
type E struct{
	gorm.Model

	Ea string `valid:"required~Field Ea is required"`
}