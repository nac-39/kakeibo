package domain

import (
	"testing"
	"time"
)

func TestNewBox(t *testing.T) {
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
		got, err := NewBox("test01", FrequencyDaily, want_duration, Active)
		if err != nil {
			t.Errorf("NewBox() error = %v", err)
		}
		if *got != *want {
			t.Errorf("NewBox() = %v, want %v", got, want)
		}

	})
	t.Run("nameが空文字", func(t *testing.T) {
		_, err := NewBox("", FrequencyDaily, Duration{}, Active)
		if err == nil {
			t.Errorf("NewBox() error = %v", err)
		}
	})
}
