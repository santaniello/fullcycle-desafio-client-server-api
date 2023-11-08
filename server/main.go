package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/santaniello/fullcycle-desafio-client-server-api/server/cotacao"
	"net/http"
)

func main() {
	mux := mux.NewRouter()
	//mux.HandleFunc("/cotacao/{cambio}", cotar)
	mux.HandleFunc("/cotacao", cotar)
	http.ListenAndServe(":8080", mux)
}

func cotar(rw http.ResponseWriter, req *http.Request) {
	/*
		vars := mux.Vars(req)
		cambioParam := vars["cambio"]
		if cambioParam == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	*/
	cotacaoInfos, err := cotacao.Cotar("USD-BRL")
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	repository, err := cotacao.NewCotacaoRepository("./cotacao.db")
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer repository.Close()

	err = repository.Save(cotacaoInfos.Bid)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(cotacao.NewCotacaoResponse(cotacaoInfos.Bid))
}
