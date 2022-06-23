package models

import "time"

type User struct {
	ID        uint64    `json: "id,omitempty"`
	Name      string    `json: "name, omitempty"`
	NickName  string    `json: "nickName, omitempty"`
	Email     string    `json:"email, omitempty"`
	Password  string    `json: "senha, omitempty"`
	CreatedAt time.Time `json: "criadoEM, omitempty"`
}
