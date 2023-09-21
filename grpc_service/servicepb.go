package grpc_service

import (
	"bytes"
	"connectrpc.com/connect"
	"context"
	"errors"
	qrpb "github.com/carolinaborim/go-qr-code-generator/proto/gen"
	qrServerpb "github.com/carolinaborim/go-qr-code-generator/proto/gen/genconnect"
	"github.com/carolinaborim/go-qr-code-generator/qr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type QRService struct{}

func (s *QRService) GenerateQr(ctx context.Context, req *connect.Request[qrpb.GenerateQrRequest]) (*connect.Response[qrpb.GenerateQrResponse], error) {
	qrurl := req.Msg.GetUrl()
	log.Println("url received", qrurl)

	if qrurl == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("missing url parameter"))
	}

	//w.Header().Set("Content-Type", "image/png")

	buffer := bytes.NewBuffer(nil)
	if err := qr.EncodeUrl(qrurl, buffer); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&qrpb.GenerateQrResponse{
		Image: buffer.Bytes(),
	}), nil
}

func RunServer() {
	mux := http.NewServeMux()
	path, handler := qrServerpb.NewQrGeneratorHandler(&QRService{})
	mux.Handle(path, handler)
	log.Println("... Listening on", "localhost:8080")
	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}

func RunGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := qrpb.RegisterQrGeneratorHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("... Listening on", "localhost:8080")
	http.ListenAndServe(":8080", mux)
}
