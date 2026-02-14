package decoder

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func ReadHexFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hexStr := strings.TrimSpace(string(data))
	raw, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("invalid hex in file")
	}

	return raw, nil
}
