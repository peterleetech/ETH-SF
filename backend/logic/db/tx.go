package db

import "gorm.io/gorm"

const (
	TxStatusUnknown = "Unknown"
	TxStatusSuccess = "Success"
	TxStatusFail    = "Fail"
)

type Tx struct {
	gorm.Model

	Name        string  `gorm:"size:200;index:idx_name;not null"`
	TaskId      string  `gorm:"size:64;not null"`
	Hash        string  `gorm:"size:200;index:idx_hash;not null"`
	Status      string  `gorm:"size:20;index:idx_name;not null"`
	Why         *string `gorm:"type:TEXT;"`
	TxData      []byte  `gorm:"type:BLOB;not null"`
	ReceiptData []byte  `gorm:"type:TEXT;"`
}

func FindTxByHashAndTaskId(hash, taskId string) (tx *Tx, err error) {
	err = db.Where("hash = ? and task_id = ?", hash, taskId).Find(&tx).Error
	return
}

func FindTxByNameAndStatus(name, status string) (txs []*Tx, err error) {
	err = db.Where("name = ? and status = ?", name, status).Find(&txs).Error
	return
}

func SaveTx(tx *Tx) (err error) {
	return db.Save(tx).Error
}
