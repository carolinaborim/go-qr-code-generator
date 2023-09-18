package qr

import (
	"github.com/carolinaborim/go-qr-code-generator/qrtst"
	"path"
	"testing"
)

func Test_EncodeUrlToFile(t *testing.T) {
	url := "https://google.com"
	outPath := path.Join(t.TempDir(), "qr-output.png")

	if err := EncodeUrlToFile(url, outPath); err != nil {
		t.Errorf("generateQRCode() error = %v", err)
	}

	if got := qrtst.ReadFile(t, outPath); got != url {
		t.Errorf("got %q want %q", got, url)
	}
}
