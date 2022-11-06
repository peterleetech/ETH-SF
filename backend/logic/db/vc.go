package db

import "gorm.io/gorm"

type VerifiableCredential struct {
	gorm.Model

	UserId uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId;references:ID"`
	VC     string `gorm:"size:4096;not null"`
}

func GetVCByUserIdAndPage(userId uint, page int, pageSize int) (vcs []*VerifiableCredential, err error) {
	query := db.
		Where("user_id = ?", userId).
		Order("id desc").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&vcs)

	err = query.Error
	return
}

func GetVCById(vcId uint) (vc *VerifiableCredential, err error) {
	query := db.
		Where("id = ?", vcId).
		Find(&vc)

	err = query.Error
	return
}

func SaveVC(vc *VerifiableCredential) (err error) {
	return db.Save(vc).Error
}
