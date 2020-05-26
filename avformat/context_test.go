package avformat

import (
	"testing"
)

func TestMain(m*testing.M) {
	// Register all formats and codecs
	AvRegisterAll()
}

func TestOpen(t*testing.T) {
	ctx := NewContext()
	err := ctx.OpenInput("small.mp4", nil, nil)
	if err != nil {
		t.Errorf("OpenInput got %q wanted nil", err)
	}
}
