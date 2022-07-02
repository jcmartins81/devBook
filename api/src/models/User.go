package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json: "id,omitempty"`
	Name      string    `json: "name, omitempty"`
	NickName  string    `json: "nickName, omitempty"`
	Email     string    `json:"email, omitempty"`
	Password  string    `json: "senha, omitempty"`
	CreatedAt time.Time `json: "criadoEM, omitempty"`
}

// Preparar vai chamar metodos para formatar e validar o usuario
func (user *User) Preparar(etapa string) error {
	if err := user.validar(etapa); err != nil {
		return err
	}
	user.formatar()
	return nil
}

func (user *User) validar(etapa string) error {
	if user.Name == "" {
		return errors.New("Nome não pode ser vazio")
	}
	if user.NickName == "" {
		return errors.New("NickName não pode ser vazio")
	}
	if user.Email == "" {
		return errors.New("Email não pode ser vazio")
	}
	if user.Password == "" {
		return errors.New("Senha não pode ser vazio")
	}

	if etapa == "cadastro" && user.Password == "" {
		return errors.New("A senha é obrigatória para realizar o cadastro")
	}

	return nil
}

func (user *User) formatar() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
}
