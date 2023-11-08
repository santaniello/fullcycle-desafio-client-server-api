package cotacao

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func NewCotacaoResponse(bid string) *CotacaoResponse {
	return &CotacaoResponse{Bid: bid}
}
