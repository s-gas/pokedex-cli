package request

import (
	"io"
	"fmt"
	"net/http"
	"errors"
)

var NotFound = errors.New("not found")

func Do(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, NotFound
		}
		return nil, fmt.Errorf("makeNewRequest: status %d", res.StatusCode)
	}
	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("makeNewRequest: %w", err)
	}
	return rawData, nil
}
