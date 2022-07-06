package main

import (
	"crypto/rand"
	"devBook/api/src/config"
	"devBook/api/src/router"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	fmt.Println(stringBase64)

	fmt.Println("teste")

}

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
