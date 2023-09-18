package qrtst

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"io"
	"os"
	"testing"
)

// ReadFile tries to decode a qr from an image file with the path given
func ReadFile(t *testing.T, path string) string {
	t.Helper()

	// open and decode image file
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file %s, %v", path, err)
	}

	return Read(t, file)
}

func Read(t *testing.T, reader io.Reader) string {
	img, _, err := image.Decode(reader)
	if err != nil {
		t.Fatalf("failed to decode image, %v", err)
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		t.Fatalf("failed to create bitmap qr, %v", err)
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		t.Fatalf("failed to decode qr, %v", err)
	}

	got := result.GetText()
	return got
}
