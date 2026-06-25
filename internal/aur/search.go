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
	PackageBase string   `json:"PackageBase"`
	Maintainer  string   `json:"Maintainer"`
	Depends     []string `json:"Depends"`
	MakeDepends []string `json:"MakeDepends"`
}

type AURResponse struct {
	Results []AURPackage `json:"results"`
}

func SearchPackages(query string) ([]AURPackage, error) {
	resp, err := http.Get(fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s", url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}
	return aurResp.Results, nil
}