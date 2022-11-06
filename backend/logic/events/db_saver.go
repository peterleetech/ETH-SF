package events

import (
	"github.com/ethereum/go-ethereum/common"
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/eth"
)

type eventDBSaver struct {
}

func (e *eventDBSaver) Save(name string, taskId string, txHash common.Hash, txBinary []byte) error {
	return db.SaveTx(&db.Tx{
		Name:   name,
		TaskId: taskId,
		Hash:   txHash.String(),
		Status: db.TxStatusUnknown,
		TxData: txBinary,
	})
}

func (e *eventDBSaver) SaveTxSuccess(taskId string, txHash common.Hash, receiptBinary []byte) error {
	tx, err := db.FindTxByHashAndTaskId(txHash.String(), taskId)
	if err != nil {
		return err
	}

	tx.Status = db.TxStatusSuccess
	tx.ReceiptData = receiptBinary
	return db.SaveTx(tx)
}

func (e *eventDBSaver) SaveTxFail(taskId string, txHash common.Hash, receiptBinary []byte, why error) error {
	tx, err := db.FindTxByHashAndTaskId(txHash.String(), taskId)
	if err != nil {
		return err
	}

	errStr := why.Error()
	tx.Status = db.TxStatusFail
	tx.Why = &errStr
	tx.ReceiptData = receiptBinary
	return db.SaveTx(tx)
}

func (e *eventDBSaver) LoadUnknownTx(name string) (unknownTx <-chan *eth.TransitionLoadItem, err error) {
	status, err := db.FindTxByNameAndStatus(name, db.TxStatusUnknown)
	if err != nil {
		return nil, err
	}

	txChan := make(chan *eth.TransitionLoadItem, 100)
	go func() {
		for _, tx := range status {
			txChan <- &eth.TransitionLoadItem{
				TaskId: tx.TaskId,
				Tx:     tx.TxData,
			}
		}
		close(txChan)
	}()

	return txChan, nil
}

func (e *eventDBSaver) SaveProcessItem(taskId string, item []byte) error {
	return db.SaveProcessItem(&db.ProcessItem{
		TaskId: taskId,
		Data:   item,
	})
}

func (e *eventDBSaver) LoadProcessItem(taskId string) (item []byte, err error) {
	dbItem, err := db.FindProcessItemByTaskId(taskId)
	if err != nil {
		return nil, err
	}

	return dbItem.Data, nil
}

func (e *eventDBSaver) HasProcessItem(taskId string) (ok bool, err error) {
	dbItem, err := db.FindProcessItemByTaskId(taskId)

	return dbItem.ID > 0, err
}

func (e *eventDBSaver) DeleteProcessItem(taskId string) error {
	return nil
}
