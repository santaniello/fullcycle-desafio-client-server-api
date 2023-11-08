package cotacao

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)


func Get(url string, duration time.Duration) (*Cotacao, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var cr Cotacao
	err = json.Unmarshal(body, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}
