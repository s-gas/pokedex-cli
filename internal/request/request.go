package request

import (
	"io"
	"fmt"
	"net/http"
)

func Do(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	defer res.Body.Close()
	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	return rawData, nil
}
