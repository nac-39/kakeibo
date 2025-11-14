package domainevent

import (
	"testing"
	"time"

	"github.com/nac-39/kakeibo/pkg/domain"
	"github.com/nac-39/kakeibo/pkg/entity"
)

func TestDepositEvent(t *testing.T) {

	t.Run("入金が正常にできる", func(t *testing.T) {
		from, err := domain.NewBox("From", domain.FrequencyNone, domain.Duration{StartDate: time.Now(), EndDate: time.Now().Add(time.Hour)}, domain.Active)
		if err != nil {
			t.Errorf("Failed to create box: %v", err)
		}
		to, err := domain.NewBox("To", domain.FrequencyNone, domain.Duration{StartDate: time.Now(), EndDate: time.Now().Add(time.Hour)}, domain.Active)
		if err != nil {
			t.Errorf("Failed to create box: %v", err)
		}

		from.Deposit(entity.Money(50))
		to.Deposit(entity.Money(0))

		_, err = NewDepositEvent(from, to, entity.Money(50))
		if err != nil {
			t.Errorf("Failed to create deposit event: %v", err)
		}
	})

	t.Run("入金に失敗する", func(t *testing.T) {
		from, err := domain.NewBox("From", domain.FrequencyNone, domain.Duration{StartDate: time.Now(), EndDate: time.Now().Add(time.Hour)}, domain.Active)
		if err != nil {
			t.Errorf("Failed to create box: %v", err)
		}
		to, err := domain.NewBox("To", domain.FrequencyNone, domain.Duration{StartDate: time.Now(), EndDate: time.Now().Add(time.Hour)}, domain.Active)
		if err != nil {
			t.Errorf("Failed to create box: %v", err)
		}

		from.Deposit(entity.Money(0))
		to.Deposit(entity.Money(0))

		event, err := NewDepositEvent(from, to, entity.Money(50))
		if err != nil {
			t.Errorf("Failed to create deposit event: %v", err)
		}

		if event.CanApply() {
			t.Errorf("DepositEvent should not be able to apply")
		}
	})

}
