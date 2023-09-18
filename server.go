package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carolinaborim/go-qr-code-generator/qr"
)

func handler(w http.ResponseWriter, req *http.Request) {
	qrurl := req.URL.Query().Get("url")

	if qrurl == "" {
		handleError(w, req, http.StatusBadRequest, map[string]string{"error": "missing url parameter"})
		return
	}

	w.Header().Set("Content-Type", "image/png")
	if err := qr.EncodeUrl(qrurl, w); err != nil {
		handleError(w, req, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
}

func handleError(w http.ResponseWriter, _ *http.Request, code int, payload map[string]string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	jsonPayload, _ := json.Marshal(payload)
	_, _ = w.Write(jsonPayload)
}

func runServer() {
	log.Println("Running server, http://localhost:8090")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
