package main

import (
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/json"
	"net/http"
)

func refreshJwks(whoamiJwksPath string) (*jose.JSONWebKeySet, error) {
	fmt.Printf("refreshing whoami jwt keys from: %s\n", whoamiJwksPath)

	resp, err := http.Get(whoamiJwksPath)
	var jwks jose.JSONWebKeySet

	if err != nil {
		return &jwks, fmt.Errorf("could not receive the new jwks config %v", err)
	}

	if resp.StatusCode != 200 {
		return &jwks, fmt.Errorf("could not receive the new jwks. The response status was %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return &jwks, fmt.Errorf("could not decode the jwts %s", err.Error())
	}

	return &jwks, nil
}

func main() {
	jwks, err := refreshJwks("https://whoami-lapi.prod.ps.welt.de/.well-known/jwks.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", jwks)
}
