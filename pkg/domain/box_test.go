package domain

import (
	"testing"
	"time"
)

func TestNewWalletBox(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		want_duration := Duration{time.Now(), time.Now().Add(time.Hour)}
		want := &Box{
			Id:           DEFAULT_ID,
			Balance:      DEFAULT_BALANCE,
			Name:         "test01",
			Frequency:    FrequencyDaily,
			Duration:     want_duration,
			Availability: Active,
		}
		got, err := NewWalletBox("test01", FrequencyDaily, want_duration, Active)
		if err != nil {
			t.Errorf("NewWalletBox() error = %v", err)
		}
		if *got != *want {
			t.Errorf("NewWalletBox() = %v, want %v", got, want)
		}

	})
	t.Run("nameが空文字", func(t *testing.T) {
		_, err := NewWalletBox("", FrequencyDaily, Duration{}, Active)
		if err == nil {
			t.Errorf("NewWalletBox() error = %v", err)
		}
	})
}
