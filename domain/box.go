package domain

import "errors"

const DEFAULT_ID = 0
const DEFAULT_BALANCE = 0

type FrequencyEnum int

const (
	FrequencyDaily FrequencyEnum = iota
	FrequencyWeekly
	FrequencyMonthly
	FrequencyYearly
	FrequencyNone
)

func (f FrequencyEnum) String() (string, error) {
	switch f {
	case FrequencyDaily:
		return "Daily", nil
	case FrequencyWeekly:
		return "Weekly", nil
	case FrequencyMonthly:
		return "Monthly", nil
	case FrequencyYearly:
		return "Yearly", nil
	case FrequencyNone:
		return "None", nil
	default:
		return "", errors.New("unknown frequency")
	}
}

type Box interface {
	Deposit(amount int) error
	Credit(amount int) error
	Transfer(amount int, to Box) error
}

type Wallet interface {
	Close() error
}

type WalletBox struct {
	id        int
	balance   int
	name      string
	frequency FrequencyEnum
}

func NewWalletBox(name string, frequency FrequencyEnum) (*WalletBox, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &WalletBox{
		id:        DEFAULT_ID,
		balance:   DEFAULT_BALANCE,
		name:      name,
		frequency: frequency,
	}, nil
}
