package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetPackageInfo(pkgName string) (*AURPackage, error) {
	resp, err := http.Get(fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=info&arg[]=%s", url.QueryEscape(pkgName)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}

	if len(aurResp.Results) == 0 {
		return nil, fmt.Errorf("package not found")
	}
	return &aurResp.Results[0], nil
}