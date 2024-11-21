package entity

type Cliente struct {
	Cpf   int64  `json:"cpf"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
