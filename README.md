[![Build Status](https://travis-ci.com/dapperlabs/dappauth.svg?branch=master)](https://travis-ci.com/dapperlabs/dappauth)
[![Coverage Status](https://coveralls.io/repos/github/dapperlabs/dappauth/badge.svg?branch=master)](https://coveralls.io/github/dapperlabs/dappauth?branch=master)
# dappauth — nugbase fork

The authors of https://flowerpatch.app manage a more up-to-date version of
dappauth / ERC1271. This lib is currently working with go-ethereum v1.9.25
(which had a bunch of annoying ABI changes and small tweaks)

## Example usage within a webserver

```Go
package main

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nugbase/dappauth"
)

// AuthenticationHandler ..
type AuthenticationHandler struct {
	client *ethclient.Client
}

// NewAuthenticationHandler ..
func NewAuthenticationHandler(rawurl string) (*AuthenticationHandler, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &AuthenticationHandler{client: client}, nil
}

// ServeHTTP serves just a single route for authentication as an example
func (a *AuthenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	challenge := r.PostFormValue("challenge")
	signature := r.PostFormValue("signature")
	addrHex := r.PostFormValue("addrHex")

	authenticator := dappauth.NewAuthenticator(r.Context(), a.client)
	isAuthorizedSigner, err := authenticator.IsAuthorizedSigner(challenge, signature, addrHex)
	if err != nil {
		// return a 5XX status code
	}
	if !isAuthorizedSigner{
		// return a 4XX status code
	}

	// create an authenticated session for address
	// return a 2XX status code
}

func main() {
	handler, err := NewAuthenticationHandler("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", handler))
}
```
