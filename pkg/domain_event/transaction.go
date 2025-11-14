package domainevent

import (
	"errors"

	"github.com/nac-39/kakeibo/pkg/domain"
	"github.com/nac-39/kakeibo/pkg/entity"
)

type TransactionEvent interface {
	CanApply() bool
}

// 貯金するお金をFromの口座からToの口座へ入金するイベント
type DepositEvent struct {
	From   *domain.Box
	To     *domain.Box
	Amount entity.Money
}

// 支出するお金をFromの口座からToの口座へ入金するイベント
type CreditEvent struct {
	From   *domain.Box
	To     *domain.Box
	Amount entity.Money
}

// 振り替えるお金をFromの口座からToの口座へ入金するイベント
type TransferEvent struct {
	From   *domain.Box
	To     *domain.Box
	Amount entity.Money
}

type ErrInsufficientBalance error

func (de DepositEvent) canApply() bool {
	return de.From.Balance >= de.Amount && de.To.Balance+de.Amount >= 0
}

func (ce CreditEvent) canApply() bool {
	return ce.From.Balance >= ce.Amount && ce.To.Balance+ce.Amount >= 0
}

func (te TransferEvent) canApply() bool {
	return te.From.Balance >= te.Amount && te.To.Balance+te.Amount >= 0
}

func NewDepositEvent(from *domain.Box, to *domain.Box, amount entity.Money) (DepositEvent, error) {
	de := DepositEvent{
		From:   from,
		To:     to,
		Amount: amount,
	}
	if !de.canApply() {
		return DepositEvent{}, errors.New("Invalid deposit event")
	}
	return de, nil
}

func NewCreditEvent(from *domain.Box, to *domain.Box, amount entity.Money) (CreditEvent, error) {
	return CreditEvent{
		From:   from,
		To:     to,
		Amount: amount,
	}, nil
}

func NewTransferEvent(from *domain.Box, to *domain.Box, amount entity.Money) (TransferEvent, error) {
	return TransferEvent{
		From:   from,
		To:     to,
		Amount: amount,
	}, nil
}
