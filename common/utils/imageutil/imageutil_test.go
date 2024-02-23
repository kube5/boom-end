package imageutil

import (
	"image/color"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func TestAddLabel(t *testing.T) {
	// Load image
	img, err := LoadImage("testdata/test.jpg")
	if err != nil {
		t.Fatalf("failed to load image: %v", err)
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		t.Fatalf("failed to parse font: %v", err)
	}

	// Add label
	img, err = AddLabel(img, "abcdeft", font, 40, color.RGBA{100, 200, 100, 200}, 100, 100)
	if err != nil {
		t.Fatalf("failed to add label: %v", err)
	}

	// Save image
	err = SaveImage(img, "testdata/test_with_label.jpg")
	if err != nil {
		t.Fatalf("failed to save image: %v", err)
	}
}
