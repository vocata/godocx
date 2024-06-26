package docxrun

import (
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
)

func TestNewRunStyle(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected *ctypes.CTString
	}{
		{
			name:     "Custom RunStyle",
			value:    "CustomStyle",
			expected: ctypes.NewCTString("CustomStyle"),
		},
		{
			name:     "Empty RunStyle",
			value:    "",
			expected: ctypes.NewCTString(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewRunStyle(tt.value)

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}

func TestDefaultRunStyle(t *testing.T) {
	expected := ctypes.NewCTString("Normal")
	result := DefaultRunStyle()

	if result.Val != expected.Val {
		t.Errorf("Expected Val %s but got %s", expected.Val, result.Val)
	}
}
