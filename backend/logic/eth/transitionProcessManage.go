package eth

import (
	"crypto/sha256"
	"encoding"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type ProcessItem interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	UniqId() string
}

type ProcessItemSaver interface {
	SaveProcessItem(taskId string, item []byte) error
	LoadProcessItem(taskId string) (item []byte, err error)
	HasProcessItem(taskId string) (ok bool, err error)
	DeleteProcessItem(taskId string) error
}

type TransitionProcessStep[T any, PT interface {
	*T
	ProcessItem
}] struct {
	stepName    string
	isFirstStep bool
	isLastStep  bool
	processFunc func(PT, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)
}

type TransitionProcessManage[T any, PT interface {
	*T
	ProcessItem
}] struct {
	tm          *TransitionManage
	saver       ProcessItemSaver
	processName string
	steps       []TransitionProcessStep[T, PT]
}

type TransitionProcessStart[T any, PT interface {
	*T
	ProcessItem
}] interface {
	FirstStep(stepName string, f func(PT, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T, PT]
	OneStep(stepName string, f func(PT) (err error)) (TransitionProcessCaller[T, PT], error)
}

type TransitionProcessMaker[T any, PT interface {
	*T
	ProcessItem
}] interface {
	NextStep(stepName string, f func(PT, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T, PT]
	LastStep(stepName string, f func(PT, *types.Transaction, *types.Receipt) (err error)) (TransitionProcessCaller[T, PT], error)
}

type TransitionProcessCaller[T any, PT interface {
	*T
	ProcessItem
}] interface {
	Run(ctx PT) error
}

func NewTransitionProcessManage[T any, PT interface {
	*T
	ProcessItem
}](tm *TransitionManage, saver ProcessItemSaver, processName string) TransitionProcessStart[T, PT] {
	return &TransitionProcessManage[T, PT]{
		tm:          tm,
		processName: processName,
		saver:       saver,
	}
}

func (m *TransitionProcessManage[T, PT]) makeFullStepName(stepName string) string {
	return fmt.Sprintf("%s::%s", m.processName, stepName)
}

func (m *TransitionProcessManage[T, PT]) addStep(stepName string, f func(PT, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) *TransitionProcessManage[T, PT] {
	m.steps = append(m.steps, TransitionProcessStep[T, PT]{
		stepName:    stepName,
		processFunc: f,
	})
	return m
}

func (m *TransitionProcessManage[T, PT]) addLastStep(stepName string, f func(PT, *types.Transaction, *types.Receipt) (err error)) (*TransitionProcessManage[T, PT], error) {
	m.addStep(stepName, func(t PT, tx *types.Transaction, r *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
		return nil, f(t, tx, r)
	})

	if len(m.steps) > 0 {
		m.steps[0].isFirstStep = true
		m.steps[len(m.steps)-1].isLastStep = true

		for i, step := range m.steps {
			err := m.tm.TransitionRegister(m.makeFullStepName(step.stepName), m.makeNextStepFunc(i))
			if err != nil {
				return nil, err
			}
		}
	}

	return m, nil
}

func (m *TransitionProcessManage[T, PT]) FirstStep(stepName string, f func(PT, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T, PT] {
	return m.addStep(stepName, func(t PT, _ *types.Transaction, _ *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
		return f(t, opts)
	})
}

func (m *TransitionProcessManage[T, PT]) OneStep(stepName string, f func(PT) (err error)) (TransitionProcessCaller[T, PT], error) {
	return m.addLastStep(stepName, func(t PT, _ *types.Transaction, _ *types.Receipt) (err error) {
		return f(t)
	})
}

func (m *TransitionProcessManage[T, PT]) NextStep(stepName string, f func(PT, *types.Transaction, *types.Receipt, *bind.TransactOpts) (next *types.Transaction, err error)) TransitionProcessMaker[T, PT] {
	return m.addStep(stepName, f)
}

func (m *TransitionProcessManage[T, PT]) LastStep(stepName string, f func(PT, *types.Transaction, *types.Receipt) (err error)) (TransitionProcessCaller[T, PT], error) {
	return m.addLastStep(stepName, f)
}

func (m *TransitionProcessManage[T, PT]) Run(ctx PT) error {
	taskId := ctx.UniqId()
	ctxBinary, err := ctx.MarshalBinary()
	if err != nil {
		return err
	}

	if len(taskId) > 50 {
		hash := sha256.Sum256([]byte(taskId))
		taskId = base64.RawStdEncoding.EncodeToString(hash[:])
	}

	hasProcessItem, err := m.saver.HasProcessItem(taskId)
	if err != nil {
		return err
	}

	if !hasProcessItem {
		err = m.saver.SaveProcessItem(taskId, ctxBinary)
		if err != nil {
			return err
		}

		if len(m.steps) > 0 {
			firstStep := m.steps[0]
			return m.tm.TransitionRequest(m.makeFullStepName(firstStep.stepName), taskId, func(taskId string, opts *bind.TransactOpts) (tx *types.Transaction, err error) {
				ctx, err := m.findCtx(taskId)
				if err != nil {
					return nil, err
				}

				return firstStep.processFunc(ctx, nil, nil, opts)
			})
		}
	}

	return nil
}

func (m *TransitionProcessManage[T, PT]) findCtx(taskId string) (ctx PT, err error) {
	item, err := m.saver.LoadProcessItem(taskId)
	if err != nil {
		return ctx, err
	}

	ctx = PT(new(T))
	err = ctx.UnmarshalBinary(item)
	return ctx, err
}

func (m *TransitionProcessManage[T, PT]) makeNextStepFunc(i int) TransitionEvent {
	return func(taskId string, tx *types.Transaction, receipt *types.Receipt) error {
		ctx, err := m.findCtx(taskId)
		if err != nil {
			return err
		}

		if receipt.Status == 0 {
			return nil
		}

		nextStep := i + 1
		if len(m.steps) > nextStep {
			step := m.steps[nextStep]
			if step.isLastStep {
				_, err = step.processFunc(ctx, tx, receipt, nil)
				if err != nil {
					return err
				}

				_ = m.saver.DeleteProcessItem(taskId)
			} else {
				return m.tm.TransitionRequest(m.makeFullStepName(step.stepName), taskId, func(taskId string, opts *bind.TransactOpts) (tx *types.Transaction, err error) {
					return step.processFunc(ctx, tx, receipt, opts)
				})
			}
		}
		return nil
	}
}
