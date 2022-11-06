package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Address string `gorm:"index:idx_address,size:50;not null"`
	DID     string `gorm:"size:4096;not null"`
	DIDId   string `gorm:"index:idx_didId,size:100;not null"`
}

func FindUserByAddress(address string) (user *User, err error) {
	err = db.Where("address = ?", address).Find(&user).Error
	return
}

func FindUserByDidID(didID string) (user *User, err error) {
	err = db.Where("did_id = ?", didID).Find(&user).Error
	return
}

func SaveUser(user *User) (err error) {
	return db.Save(user).Error
}
