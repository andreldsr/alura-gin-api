package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name" validate:"min=3,regexp=^[0-9]*$"`
	Document string `json:"document" validate:"min=9,max=11,regexp=^[0-9]*$"`
}

func ValidateStudent(student Student) (err error) {
	err = validator.Validate(student)
	return
}
