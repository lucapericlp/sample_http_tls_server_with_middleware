package main

import (
	"./home"
	"./server"
	"log"
	"net/http"
	"os"
)

var (
	//generate using following commands:
	/*
		openssl req -x509 -out localhost.crt -keyout localhost.key \
		  -newkey rsa:2048 -nodes -sha256 \
		  -subj '/CN=localhost' -extensions EXT -config <( \
		   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
	*/
	CertFile    = os.Getenv("TEST_CERT_FILE")
	KeyFile     = os.Getenv("TEST_KEY_FILE")
	ServiceAddr = os.Getenv("TEST_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "test", log.LstdFlags|log.Lshortfile)
	h := home.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	srv := server.New(mux, ServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
