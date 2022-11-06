package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
	"time"
	"veric-backend-mvp/logic/log"
	"veric-backend-mvp/logic/util"
)

type TransitionQueueItem struct {
	name   string
	taskId string
	tx     *types.Transaction
}

type TransitionLoadItem struct {
	TaskId string
	Tx     []byte
}

type TransitionSaver interface {
	Save(name string, taskId string, txHash common.Hash, txBinary []byte) error
	SaveTxSuccess(taskId string, txHash common.Hash, receiptBinary []byte) error
	SaveTxFail(taskId string, txHash common.Hash, receiptBinary []byte, why error) error
	LoadUnknownTx(name string) (unknownTx <-chan *TransitionLoadItem, err error)
}

type TransitionEvent func(taskId string, tx *types.Transaction, receipt *types.Receipt) error
type TransitionRequest func(taskId string, opts *bind.TransactOpts) (tx *types.Transaction, err error)

type TransitionManage struct {
	eventsMap  *util.SyncedMap[string, TransitionEvent]
	account    *PrivateKey
	chainId    *big.Int
	ethClient  *Client
	saver      TransitionSaver
	checkQueue chan *TransitionQueueItem
}

func NewTransitionManage(account *PrivateKey, chainId *big.Int, ethClient *Client, saver TransitionSaver) *TransitionManage {
	manager := &TransitionManage{
		eventsMap:  new(util.SyncedMap[string, TransitionEvent]),
		account:    account,
		chainId:    chainId,
		ethClient:  ethClient,
		saver:      saver,
		checkQueue: make(chan *TransitionQueueItem, 1000),
	}
	manager.init()

	return manager
}

func (m *TransitionManage) init() {
	go m.checkTransition()
}

func (m *TransitionManage) AddTransitionCheckQueue(name string, taskId string, tx *types.Transaction) func() {
	return func() {
		m.checkQueue <- &TransitionQueueItem{
			name:   name,
			taskId: taskId,
			tx:     tx,
		}
	}
}

func (m *TransitionManage) TransitionRegister(name string, te TransitionEvent) error {
	m.eventsMap.Store(name, te)
	return m.loadUnCheckedTransition(name)
}

func (m *TransitionManage) TransitionRequest(name, taskId string, te TransitionRequest) error {
	transact, err := m.account.MakeTransact(m.chainId)
	if err != nil {
		return err
	}

	nonce, err := m.ethClient.Nonce(context.Background(), m.account.Address())
	if err != nil {
		return err
	}

	transact.Nonce = big.NewInt(int64(nonce))
	transact.Value = big.NewInt(0)
	transact.GasLimit = uint64(3000000)

	transaction, err := te(taskId, transact)
	if err != nil {
		return err
	}

	txBinary, err := transaction.MarshalBinary()
	if err != nil {
		return err
	}

	time.AfterFunc(5*time.Second, m.AddTransitionCheckQueue(name, taskId, transaction))

	return m.saver.Save(name, taskId, transaction.Hash(), txBinary)
}

func (m *TransitionManage) checkTransition() {
	for item := range m.checkQueue {
		receipt, err := m.ethClient.TransactionReceipt(item.tx.Hash())
		if err != nil {
			time.AfterFunc(15*time.Second, m.AddTransitionCheckQueue(item.name, item.taskId, item.tx))
			log.GetLogger().Warn("TransactionReceipt check error", zap.Error(err))
			continue
		}

		receiptBinary, err := receipt.MarshalJSON()
		if err != nil {
			log.GetLogger().Warn("TransactionReceipt MarshalJSON error", zap.Error(err))
			continue
		}

		if eventFunc, ok := m.eventsMap.Load(item.name); ok {
			err = eventFunc(item.taskId, item.tx, receipt)
			if err != nil {
				err = m.saver.SaveTxFail(item.taskId, item.tx.Hash(), receiptBinary, err)
				if err != nil {
					log.GetLogger().Warn("transaction save error", zap.Error(err))
					continue
				}
			}

			if receipt.Status == 1 {
				err = m.saver.SaveTxSuccess(item.taskId, item.tx.Hash(), receiptBinary)
			} else {
				err = m.saver.SaveTxFail(item.taskId, item.tx.Hash(), receiptBinary, errors.New("receipt status is 0"))
			}

			if err != nil {
				log.GetLogger().Warn("transaction save error", zap.Error(err))
				continue
			}
		}
	}
}

func (m *TransitionManage) loadUnCheckedTransition(name string) error {
	unknownChan, err := m.saver.LoadUnknownTx(name)
	if err != nil {
		return err
	}

	for unknownTx := range unknownChan {
		tx := &types.Transaction{}
		err = tx.UnmarshalBinary(unknownTx.Tx)
		if err != nil {
			return err
		}

		m.checkQueue <- &TransitionQueueItem{
			name:   name,
			taskId: unknownTx.TaskId,
			tx:     tx,
		}
	}

	return nil
}
