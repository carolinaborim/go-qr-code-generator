package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"os"
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
}
