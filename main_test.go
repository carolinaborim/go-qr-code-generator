package main

import (
	"github.com/carolinaborim/go-qr-code-generator/qrtst"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	outPath := "/tmp/qr-output.png"
	url := "https://google.com"

	os.Args = []string{"test-app", "-u", url, "-o", outPath}

	main()

	if _, err := os.Stat(outPath); err != nil {
		t.Errorf("Expected file %q to exist", outPath)
	}

	got := qrtst.ReadFile(t, outPath)

	if got != url {
		t.Errorf("got %q want %q", got, url)
	}
}

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    ParsedArgs
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			wantErr: true,
		},
		{
			name: "flag with =",
			args: []string{"app-name", "-url=http://example.com", "-output=/tmp/output"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags single dash",
			args: []string{"app-name", "-url", "http://example.com", "-output", "/tmp/output"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags double dash",
			args: []string{"app-name", "--url", "http://example.com", "--output", "/tmp/output"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags double dash different order",
			args: []string{"app-name", "--output", "/tmp/output", "--url", "http://example.com"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags short names",
			args: []string{"app-name", "-o", "/tmp/output", "-u", "http://example.com"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags short and long names",
			args: []string{"app-name", "-o", "/tmp/output", "--output", "/tmp/output", "-u", "http://example.com"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "flags short and long names",
			args: []string{"app-name", "-o", "/tmp/output", "--output", "/tmp/output-long", "-u", "http://example.com"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output-long",
			},
		},
		{
			name: "flags short and long names",
			args: []string{"app-name", "--output", "/tmp/output-long", "-o", "/tmp/output", "-u", "http://example.com"},
			want: ParsedArgs{
				Url:        "http://example.com",
				OutputPath: "/tmp/output",
			},
		},
		{
			name: "with server mode",
			args: []string{"app-name", "-server"},
			want: ParsedArgs{
				ServerMode: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
