package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Name     string `gorm: "not null"`
	Birthday time.Time
	Age      int    `grom:"-"` // ignore
	gorm.Model
}

func (s *User) calAge()  {
	today := time.Now()
	s.Age = today.Year() - s.Birthday.Year()
}
