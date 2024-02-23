package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype"
)

// LoadImage loads an image file at the specified path and returns an image.Image.
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image file: %v", err)
	}

	switch format {
	case "jpeg":
		fallthrough
	case "jpg":
		return img, nil
	case "png":
		return RemoveTransparency(img), nil
	default:
		return nil, fmt.Errorf("unsupported image format: %v", format)
	}
}

// Resize resizes an image to the specified width and height.
func Resize(img image.Image, width, height int) image.Image {
	return resize(img, width, height)
}

func resize(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	ratio := bounds.Dx() / bounds.Dy()
	newHeight := height
	newWidth := newHeight * ratio

	if newWidth > width {
		newWidth = width
		newHeight = newWidth / ratio
	}

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.Draw(newImg, newImg.Bounds(), img, bounds.Min, draw.Src)
	return newImg
}

// RemoveTransparency converts a transparent image to a white background image.
func RemoveTransparency(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	draw.Draw(newImg, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Over)

	return newImg
}

// SaveImage saves an image to a file at the specified path.
func SaveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create image file: %v", err)
	}
	defer file.Close()

	switch ext := getFileExtension(path); ext {
	case "jpg":
		return jpeg.Encode(file, img, nil)
	case "png":
		return png.Encode(file, img)
	default:
		return fmt.Errorf("unsupported image format: %v", ext)
	}
}

// getFileExtension returns the extension of a file.
func getFileExtension(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}

	return ""
}

// AddLabel adds a label to an image with the specified text, font file, font size, and text color.
func AddLabel(img image.Image, text string, font0 *truetype.Font, fontSize float64, color color.Color, x, y int) (image.Image, error) {
	rgba := ToRGBA(img)

	face := truetype.NewFace(font0, &truetype.Options{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	drawer := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(color),
		Face: face,
	}

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}

	drawer.DrawString(text)

	return rgba, nil
}

// AddLabel adds a label to an image with the specified text, font file, font size, and text color.
func AddLabel2(img image.Image, label, fontFile string, fontSize float64, textColor color.Color) (image.Image, error) {
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %v", err)
	}

	bounds := img.Bounds()
	dpi := float64(72)
	fg := image.NewUniform(textColor)
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, img, image.Point{}, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(fg)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(fontSize)>>6))
	_, err = c.DrawString(label, pt)
	if err != nil {
		return nil, fmt.Errorf("failed to add label to image: %v", err)
	}

	return dst, nil
}

func ToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
