package cotacao

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const url = "https://economia.awesomeapi.com.br/json/last/"
const timeout = 200 * time.Millisecond

func Cotar(cambio string) (*Cotacao, error) {
	res, err := doRequestWithContext(cambio)
	if err != nil {
		return nil, err
	}
	cotacao, err := readResponseBody(cambio, res)
	if err != nil {
		return nil, err
	}
	return cotacao, nil
}

func doRequestWithContext(cambio string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url+cambio, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Erro ao chamar a api de cambio,  Motivo: %s", err)
		return nil, err
	}
	return res, nil
}

func readResponseBody(cambio string, res *http.Response) (*Cotacao, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]Cotacao

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	cambioFormatado := strings.ReplaceAll(cambio, "-", "")
	var cotacao = data[cambioFormatado]
	return &cotacao, nil
}
