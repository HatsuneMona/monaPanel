package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

type CommonTime struct {
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type DeleteStat struct {
	deleteStat gorm.DeletedAt `json:"delete_stat" gorm:"index"`
}
