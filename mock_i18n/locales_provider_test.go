package mock_i18n

import (
	"github.com/strongo/i18n"
	"go.uber.org/mock/gomock"
	"testing"
)

var _ i18n.LocalesProvider = (*MockLocalesProvider)(nil)

func TestNewMockLocalesProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	if NewMockLocalesProvider(ctrl) == nil {
		t.Fatal("localesProvider is nil")
	}
}
