package controllers

import (
	"devBook/api/src/banco"
	"devBook/api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(error)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro€ý,€ý,)
	}
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todos os usuÃ¡rios"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca um UsuÃ¡rio"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario"))
}
