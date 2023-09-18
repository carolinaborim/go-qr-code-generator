package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/carolinaborim/go-qr-code-generator/qr"
)

func main() {
	args, err := parseArgs(os.Args)
	if err != nil {
		log.Fatalf("Parse args failed, %v", err)
	}

	if args.ServerMode {
		runServer()
		return
	}

	if err := qr.EncodeUrlToFile(args.Url, args.OutputPath); err != nil {
		log.Fatalf("Generating qr failed, %v", err)
	}
}

type ParsedArgs struct {
	ServerMode bool
	Url        string
	OutputPath string
}

func parseArgs(args []string) (ParsedArgs, error) {
	var url string
	var output string
	var serverMode bool

	if len(args) < 1 {
		return ParsedArgs{}, fmt.Errorf("no args were passed")
	}

	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cli.SetOutput(os.Stdout)
	cli.StringVar(&url, "url", "", "url for qr-code")
	cli.StringVar(&url, "u", "", "url for qr-code")

	cli.StringVar(&output, "output", "", "output file")
	cli.StringVar(&output, "o", "", "output file")

	cli.BoolVar(&serverMode, "server", false, "run qr cli in server mode")

	err := cli.Parse(args[1:])

	return ParsedArgs{ServerMode: serverMode, Url: url, OutputPath: output}, err
}
