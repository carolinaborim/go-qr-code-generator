package main

import (
	"flag"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
)

func main() {
	args, err := parseArgs(os.Args)
	if err != nil {
		log.Fatalf("Parse args failed, %v", err)
	}

	if err := generateQRCode(args.Url, args.OutputPath); err != nil {
		log.Fatalf("Generating qr failed, %v", err)
	}
}

type ParsedArgs struct {
	Url        string
	OutputPath string
}

func parseArgs(args []string) (ParsedArgs, error) {
	var url string
	var output string

	if len(args) < 1 {
		return ParsedArgs{}, fmt.Errorf("no args were passed")
	}

	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cli.SetOutput(os.Stdout)
	cli.StringVar(&url, "url", "", "url for qr-code")
	cli.StringVar(&url, "u", "", "url for qr-code")

	cli.StringVar(&output, "output", "", "output file")
	cli.StringVar(&output, "o", "", "output file")
	err := cli.Parse(args[1:])

	return ParsedArgs{Url: url, OutputPath: output}, err
}

func generateQRCode(url string, output string) error {
	return qrcode.WriteFile(url, qrcode.Medium, 256, output)
}
