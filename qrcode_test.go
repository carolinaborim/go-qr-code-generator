package main

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"os"
	"testing"
)

func TestMainFunc_ArgParsing(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		shouldPanic bool
	}{
		{
			name:        "no args",
			args:        []string{},
			shouldPanic: true,
		},
		{
			name:        "correct agrs",
			args:        []string{"app-name", "url", "/tmp/output"},
			shouldPanic: false,
		},
		{
			name:        "correct agrs",
			args:        []string{"app-name", "--url=http://example.com", "--output=/tmp/output"},
			shouldPanic: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			defer func() {
				if err := recover(); err != nil {
					if !tt.shouldPanic {
						t.Errorf("Code paninced but it shouldn't, %v", err)
					}
				} else if tt.shouldPanic {
					t.Error("Code did not panic but it should")
				}
			}()

			os.Args = tt.args
			main()
		})
	}
}

func TestMainFunc(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	outPath := "/tmp/qr-output.png"
	url := "https://google.com"

	os.Args = []string{"test-app", url, outPath}

	main()

	if _, err := os.Stat(outPath); err != nil {
		t.Errorf("Expected file %q to exist", outPath)
	}

	got := readsQRFile(t, outPath)

	if got != url {
		t.Errorf("got %q want %q", got, url)
	}
}

func readsQRFile(t *testing.T, outPath string) string {
	t.Helper()

	// open and decode image file
	file, err := os.Open(outPath)
	if err != nil {
		t.Fatalf("failed to open file %s, %v", outPath, err)
	}
	img, _, err := image.Decode(file)
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

func TestExample(t *testing.T) {
	type Something struct{ value int }
	s := []Something{{1}, {2}, {3}, {4}}

	//ss := make([]Something, 2)
	//copy(ss, s[1:3])

	//ss := s[1:3]

	ss := append([]Something{}, s...)
	ss[0].value = 5
	t.Log(s)
	t.Log(ss)
	//
	//mySlice := []string{"a", "b", "c", "d", "e"}
	//t.Log(mySlice[1:3]) // "b", "c"
	//t.Log(mySlice[1:])  //"b", "c", "d", "e"
	//t.Log(mySlice[:3])  //"a", "b", "c"
}

func Test_generateQRCode(t *testing.T) {
	outPath := "/tmp/qr-output.png"
	url := "https://google.com"
	if err := generateQRCode(url, outPath); err != nil {
		t.Errorf("generateQRCode() error = %v", err)
	}

	//if _, err := os.Stat(outPath); err != nil {
	//	t.Errorf("Expected file %q to exist", outPath)
	//}

	got := readsQRFile(t, outPath)

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
