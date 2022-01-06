package sitemodels

import (
	"time"

	"gorm.io/gorm"
)

type AboutUs struct {
	Id          uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ImageUrl    string         `json:"image_url"`
	Content     string         `json:"content"`
	Description string         `json:"description"`
	WhyUs       string         `json:"why_us"`
	HowWorking  string         `json:"how_working"`
}