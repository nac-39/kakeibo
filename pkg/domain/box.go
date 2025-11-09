package domain

import (
	"errors"
	"time"

	"github.com/nac-39/kakeibo/pkg/entity"
)

const DEFAULT_ID = 0
const DEFAULT_BALANCE = entity.Money(0)

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

type Duration struct {
	StartDate time.Time
	EndDate   time.Time
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

type BoxName string

func (b BoxName) IsValid() bool {
	if len(b) == 0 {
		return false
	}
	return true
}

type Box struct {
	Id           int
	Balance      entity.Money
	Name         BoxName
	Frequency    FrequencyEnum
	Duration     Duration
	Availability AvailabilityEnum
}

func NewBox(name BoxName, frequency FrequencyEnum, duration Duration, availability AvailabilityEnum) (*Box, error) {
	if !name.IsValid() {
		return nil, errors.New("invalid name")
	}
	if !duration.IsValid() {
		return nil, errors.New("invalid duration")
	}
	if availability != Active {
		return nil, errors.New("new wallet box must be active")
	}

	return &Box{
		Id:           DEFAULT_ID,
		Balance:      DEFAULT_BALANCE,
		Name:         name,
		Frequency:    frequency,
		Duration:     duration,
		Availability: availability,
	}, nil
}
