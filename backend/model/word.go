package model

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	UserID     uint       `gorm:"not null;index" json:"user_id"`
	Word       string     `gorm:"type:varchar(128);not null" json:"word"`
	Definition string     `gorm:"type:text;not null" json:"definition"`
	AIProvider string     `gorm:"column:ai_provider;type:varchar(32);not null" json:"ai_provider"`
	Sentences  []Sentence `gorm:"foreignKey:WordID" json:"sentences"`
}

type Sentence struct {
	gorm.Model
	WordID  uint   `gorm:"not null;index" json:"word_id"`
	English string `gorm:"type:text;not null" json:"english"`
	Chinese string `gorm:"type:text;not null" json:"chinese"`
}
