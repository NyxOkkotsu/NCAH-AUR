package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AURPackage struct {
	Name        string   `json:"Name"`
	Version     string   `json:"Version"`
	Description string   `json:"Description"`
	URL         string   `json:"URL"`
	Maintainer  string   `json:"Maintainer"`
	PackageBase string   `json:"PackageBase"`
	Depends     []string `json:"Depends"`
	MakeDepends []string `json:"MakeDepends"`
	OptDepends  []string `json:"OptDepends"`
	License     []string `json:"License"`
	NumVotes    int      `json:"NumVotes"`
	Popularity  float64  `json:"Popularity"`
}

type RPCResponse struct {
	Results []AURPackage `json:"results"`
}

func SearchPackages(query string) ([]AURPackage, error) {
	resp, err := http.Get(fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s", url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Results, nil
}

func GetPackageInfo(name string) (*AURPackage, error) {
	resp, err := http.Get(fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=info&arg[]=%s", url.QueryEscape(name)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if len(r.Results) == 0 {
		return nil, fmt.Errorf("package not found")
	}
	return &r.Results[0], nil
}