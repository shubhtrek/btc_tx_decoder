package decoder

import (
	"fmt"
	"io"
	"net/http"
)

// Fetch tx hex from mempool.space
func FetchTxHex(txid string) (string, error) {
	url := "https://mempool.space/api/tx/" + txid + "/hex"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("network error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

