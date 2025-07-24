package mock_i18n

import (
	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
	"testing"
)

var _ i18n.Translator = (*MockTranslator)(nil)

func TestTranslator(t *testing.T) {
	ctrl := gomock.NewController(t)
	if NewMockTranslator(ctrl) == nil {
		t.Fatal("translator is nil")
	}
}
