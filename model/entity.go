package model

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ClientId     string `gorm:"index;unique;not-null"`
	ClientSecret string `gorm:"size:100"`
	Name         string `gorm:""`
	RedirectUri  string `gorm:""`
	UserID       uint
	User         User `gorm:"references:ID"`
}

func (a *Application) ToDTO() *ApplicationDto {
	return &ApplicationDto{
		ClientId:    a.ClientId,
		Name:        a.Name,
		RedirectURL: a.RedirectUri,
	}
}
