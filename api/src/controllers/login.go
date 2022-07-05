package controllers

import (
	"devBook/api/src/autenticacao"
	"devBook/api/src/banco"
	"devBook/api/src/models"
	"devBook/api/src/repositorios"
	"devBook/api/src/respostas"
	"devBook/api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Login é responsável por autenticar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.User
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificaSenha(usuarioSalvoNoBanco.Password, usuario.Password); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	fmt.Println(token)

}
