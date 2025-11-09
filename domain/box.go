package domain

import (
	"errors"
	"time"
)

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

type AvailabilityEnum int

const (
	Active AvailabilityEnum = iota
	Inactive
	Deleted
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
	Balance() int
}

type Duration struct {
	StartDate time.Time
	EndDate   time.Time
}
type WalletBox struct {
	Id           int
	Balance      int
	Name         string
	Frequency    FrequencyEnum
	Duration     Duration
	Availability AvailabilityEnum
}

func (d Duration) IsValid() bool {
	if d.StartDate.IsZero() || d.EndDate.IsZero() {
		return false
	}
	if d.EndDate.Before(d.StartDate) {
		return false
	}
	return true
}

func NewWalletBox(name string, frequency FrequencyEnum, duration Duration, availability AvailabilityEnum) (*WalletBox, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !duration.IsValid() {
		return nil, errors.New("invalid duration")
	}
	if availability != Active {
		return nil, errors.New("new wallet box must be active")
	}

	return &WalletBox{
		Id:           DEFAULT_ID,
		Balance:      DEFAULT_BALANCE,
		Name:         name,
		Frequency:    frequency,
		Duration:     duration,
		Availability: availability,
	}, nil
}
