package events

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
	"time"
	"veric-backend-mvp/logic/config"
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/eth"
	"veric-backend-mvp/logic/log"
	"veric-backend-mvp/logic/model"
	"veric-backend-mvp/logic/sol/ledger"
	"veric-backend-mvp/logic/sol/vault"
)

const maxFetchLogsOnce = 1000

type Events struct {
	conf *config.RpcConfig

	account   *eth.PrivateKey
	ethClient *eth.Client
	vaultSol  *vault.Sol
	ledgerSol *ledger.Sol

	depositedProcess eth.TransitionProcessCaller[EventDeposited, *EventDeposited]
	withdrawnProcess eth.TransitionProcessCaller[EventWithdrawn, *EventWithdrawn]
}

func NewEvents(conf *config.RpcConfig) (*Events, error) {
	e := &Events{
		conf: conf,
	}

	err := e.initClient()
	if err != nil {
		return nil, err
	}

	err = e.initProcess()
	if err != nil {
		return nil, err
	}

	err = e.initListen()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Events) initClient() error {
	var err error

	e.ethClient, err = eth.NewClient(e.conf.RPCHost)
	if err != nil {
		return err
	}

	e.vaultSol, err = e.ethClient.GetVaultSol(e.conf.VaultContractAddress)
	if err != nil {
		return err
	}

	e.ledgerSol, err = e.ethClient.GetLedgerSol(e.conf.LedgerContractAddress)
	if err != nil {
		return err
	}

	e.account, err = eth.NewPrivateKey(e.conf.AccountPriKey)
	if err != nil {
		return err
	}

	return nil
}

func (e *Events) initListen() error {
	currentNumber, err := e.ethClient.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	{
		depositedChan := make(chan *vault.SolDeposited, 100)
		go e.filterRange(currentNumber-10000, func(watchStart, watchEnd uint64) error {
			deposited, err := e.vaultSol.FilterDeposited(&bind.FilterOpts{Start: watchStart, End: &watchEnd}, nil, nil, nil)
			if err != nil {
				return err
			}
			defer deposited.Close()

			for deposited.Next() {
				if deposited.Error() != nil {
					return deposited.Error()
				}
				depositedChan <- deposited.Event
			}

			return nil
		})

		go e.OnDeposited(depositedChan)
	}

	return nil
}

func (e *Events) filterRange(watchStart uint64, f func(watchStart, watchEnd uint64) error) {
	for range time.Tick(5 * time.Second) {
		watchEnd, err := e.ethClient.BlockNumber(context.Background())
		if err != nil {
			log.GetLogger().Warn("filterRange get BlockNumber err", zap.Error(err))
			continue
		}

		for watchStart < watchEnd {
			watchTo := watchStart + maxFetchLogsOnce
			if watchTo > watchEnd {
				watchTo = watchEnd
			}

			log.GetLogger().Info("watch range changed", zap.Uint64("watchStart", watchStart), zap.Uint64("watchTo", watchTo))
			err = f(watchStart, watchTo)
			if err != nil {
				log.GetLogger().Warn("filterRange err", zap.Error(err))
				continue
			}

			watchStart = watchTo + 1
		}
	}
}

func (e *Events) initProcess() error {
	saver := &eventDBSaver{}
	manage := eth.NewTransitionManage(e.account, big.NewInt(e.conf.ChainId), e.ethClient, saver)

	{
		e.depositedProcess, _ = eth.NewTransitionProcessManage[EventDeposited](manage, saver, "EventDeposited").
			FirstStep("AddPaymentRequest", func(i *EventDeposited, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				return e.ledgerSol.AddPaymentRequest(opts, i.Deposited.TokenAddress.String(), i.Deposited.Amount, i.Deposited.Raw.TxHash.String(), i.Deposited.Who.String())
			}).
			NextStep("ConfirmPaymentRequest", func(i *EventDeposited, tx *types.Transaction, receipt *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				return e.ledgerSol.ConfirmPaymentRequest(opts, i.Deposited.Raw.TxHash.String(), i.Deposited.Who.String())
			}).
			LastStep("Result", func(i *EventDeposited, tx *types.Transaction, receipt *types.Receipt) (err error) {
				user, err := db.FindUserByAddress(i.Deposited.Who.String())
				if err != nil {
					return err
				}

				if user.ID == 0 {
					user = &db.User{
						Address: i.Deposited.Who.String(),
					}
					err = db.SaveUser(user)
					if err != nil {
						return err
					}
				}

				subject := &model.VCSubjectDeposit{
					TokenAddress: i.Deposited.TokenAddress,
					Amount:       i.Deposited.Amount,
					Vault:        e.conf.VaultContractAddress,
				}

				vcModel := model.CreateVerifiableCredential(config.IssueDID, subject)
				err = vcModel.Signature(config.IssuePriKey)
				if err != nil {
					return err
				}

				vc := &db.VerifiableCredential{
					UserId: user.ID,
					VC:     string(vcModel.ToJson()),
				}

				return db.SaveVC(vc)
			})
	}

	{
		e.withdrawnProcess, _ = eth.NewTransitionProcessManage[EventWithdrawn](manage, saver, "EventWithdrawn").
			FirstStep("AddWithdrawRequest", func(i *EventWithdrawn, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				return e.ledgerSol.AddWithdrawRequest(opts, i.TokenAddress.String(), i.Amount, i.VCId, i.Vault, i.Requester.String())
			}).
			NextStep("Withdraw", func(i *EventWithdrawn, tx *types.Transaction, receipt *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				return e.vaultSol.Withdraw(opts, i.TokenAddress, i.Requester, i.Amount)
			}).
			NextStep("ConfirmWithdrawRequest", func(i *EventWithdrawn, tx *types.Transaction, receipt *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				return e.ledgerSol.ConfirmWithdrawRequest(opts, i.VCId, i.Requester.String())
			}).
			LastStep("Result", func(i *EventWithdrawn, tx *types.Transaction, receipt *types.Receipt) (err error) {
				return nil
			})
	}

	return nil
}

func (e *Events) OnDeposited(depositedChan <-chan *vault.SolDeposited) {
	var err error
	for deposited := range depositedChan {
		log.GetLogger().Info("found deposited", zap.Any("deposited", deposited))
		err = e.depositedProcess.Run(&EventDeposited{
			Deposited: deposited,
		})
		if err != nil {
			log.GetLogger().Warn("OnDeposited error", zap.Error(err))
		}
	}
}

func (e *Events) Close() {

}

func (e *Events) SubmitWithdrawn(ev *EventWithdrawn) error {
	return e.withdrawnProcess.Run(ev)
}
