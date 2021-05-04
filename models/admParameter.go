package models

type AdmParameter struct {
	Id          uint64 `gorm:"column:par_seq;primaryKey;autoIncrement" json:"id"`
	Code        string `gorm:"column:par_code;not null" json:"code"`
	Description string `gorm:"column:par_description;not null;unique" json:"description"`
	//idParameterCategory uint64               `gorm:"column:par_pmc_seq"`
	Value             string               `gorm:"column:par_value" json:"value"`
	ParameterCategory AdmParameterCategory `gorm:"embedded" json:"admParameterCategory"`
}
