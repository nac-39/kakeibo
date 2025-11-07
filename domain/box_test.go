package domain

import (
	"testing"
)

func TestNewWalletBox(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		want := &WalletBox{
			id:        DEFAULT_ID,
			balance:   DEFAULT_BALANCE,
			name:      "test01",
			frequency: FrequencyDaily,
		}
		got, err := NewWalletBox("test01", FrequencyDaily)
		if err != nil {
			t.Errorf("NewWalletBox() error = %v", err)
		}
		if *got != *want {
			t.Errorf("NewWalletBox() = %v, want %v", got, want)
		}

	})
	t.Run("nameが空文字", func(t *testing.T) {
		_, err := NewWalletBox("", FrequencyDaily)
		if err == nil {
			t.Errorf("NewWalletBox() error = %v", err)
		}
	})
}
