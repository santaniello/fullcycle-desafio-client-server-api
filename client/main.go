package main

import (
	"fmt"
	"github.com/santaniello/fullcycle-desafio-client-server-api/client/cotacao"
	"os"
	"time"
)

func main() {
	resp, err := cotacao.Get("http://localhost:8080/cotacao", 300*time.Millisecond)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Bid)
	file := createFile()
	defer file.Close()
	writeFile(file, "Dólar: "+resp.Bid)

}

func writeFile(file *os.File, message string) int {
	tamanho, err := file.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println(err)
	}
	return tamanho
}

func createFile() *os.File {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println(err)
	}
	return file
}
