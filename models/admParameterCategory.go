package models

import "github.com/jinzhu/gorm"

type AdmParameterCategory struct {
	gorm.Model
	Description string `json:"description"`
}
