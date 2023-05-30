package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// load CA certificate file and add it to list of client CAs
	cert, err := ioutil.ReadFile("/home/appscodepc/go/src/github.com/AshrafulHaqueToni/mTLS/certs/ca.crt")
	if err != nil {
		log.Fatalf("couldn't open ca.cart file")
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	client := http.Client{
		Timeout: time.Minute * 1,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
	}
	resp, err := client.Get("https://localhost:9090")
	if err != nil {
		log.Fatalf("error making get request: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}
	fmt.Println(string(body))

}
