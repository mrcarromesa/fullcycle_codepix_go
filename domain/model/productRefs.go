package model

import "gorm.io/gorm"

type ProductRefs struct {
	gorm.Model
	Code  string
	Price uint
}
