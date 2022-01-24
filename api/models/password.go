package models

import "golang.org/x/crypto/bcrypt"

type Password struct {
	PlainText string
	Hash      []byte
}

func (p *Password) Set() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.PlainText), bcrypt.MinCost)
	if err != nil {
		return err
	}
	p.Hash = hash
	return nil
}

func (p *Password) Check(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}
