package mock_i18n

import (
	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
	"testing"
)

var _ i18n.SingleLocaleTranslator = (*MockSingleLocaleTranslator)(nil)

func TestNewMockSingleLocaleTranslator(t *testing.T) {
	ctrl := gomock.NewController(t)
	if NewMockSingleLocaleTranslator(ctrl) == nil {
		t.Fatal("translator is nil")
	}
}
