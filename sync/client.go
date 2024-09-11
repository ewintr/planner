package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	url    string
	apiKey string
	c      *http.Client
}

func NewClient(url, apiKey, certFile string) (*Client, error) {
	caCert, err := os.ReadFile(certFile)
	if err != nil {
		return nil, fmt.Errorf("could not read cert file: %s", err)
	}

	// Create a certificate pool and add the server's certificate
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a custom TLS configuration
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Create a custom transport using the TLS configuration
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Create a client using the custom transport
	c := &http.Client{
		Transport: transport,
	}

	return &Client{
		url:    url,
		apiKey: apiKey,
		c:      c,
	}, nil
}
