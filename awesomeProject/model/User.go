package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password Password
}
type Password struct {
	plaintext *string
	hash      []byte
}

func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}
