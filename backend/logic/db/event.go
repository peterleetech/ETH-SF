package db

import "gorm.io/gorm"

type Event struct {
	gorm.Model

	EventType    string `gorm:"index:idx_eventType;not null"`
	EventContent string `gorm:"size:4096;not null"`
}

func GetEventByPage(page int, pageSize int) (evs []*Event, err error) {
	query := db.
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&evs)

	err = query.Error
	return
}
