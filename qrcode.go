package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"os"
	//"github.com/makiuchi-d/gozxing"
	//"github.com/makiuchi-d/gozxing/oned"
)
import "fmt"

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 3 {
		panic("need url and file output")
	}

	url := os.Args[1]
	output := os.Args[2]

	err := qrcode.WriteFile(url, qrcode.Medium, 256, output)
	if err != nil {
		panic(err)
	}

	// Another library that will be used for reading QR codes.
	//enc := oned.NewCode128Writer()
	//img, _ := enc.Encode(url, gozxing.BarcodeFormat_QR_CODE BarcodeFormat_CODE_128, 250, 50, nil)
	//
	//file, _ := os.Create(output)
	//defer file.Close()
	//
	//// *BitMatrix implements the image.Image interface,
	//// so it is able to be passed to png.Encode directly.
	//_ = png.Encode(file, img)
}
