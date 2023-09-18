package qr

import (
	"github.com/skip2/go-qrcode"
	"io"
)

func EncodeUrl(url string, w io.Writer) error {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	_, err = w.Write(png)
	return err
}

func EncodeUrlToFile(url, filePath string) error {
	//f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 400)
	//if err != nil {
	//	return err
	//}
	//
	//return EncodeUrl(url, f)
	return qrcode.WriteFile(url, qrcode.Medium, 256, filePath)
}
