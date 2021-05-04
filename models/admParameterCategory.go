package models

type AdmParameterCategory struct {
	//gorm.Model
	Id          uint64 `gorm:"column:pmc_seq;primaryKey;autoIncrement" json:"id"`
	Description string `gorm:"column:pmc_description;not null;unique" json:"description"`
	Order       uint64 `gorm:"column:pmc_order" json:"order"`
}
