package models

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserModel struct {
	graphDb *GraphDb
}

func NewUserModel(gdb *GraphDb) *UserModel {
	return &UserModel{
		graphDb: gdb}
}

// In node I will never pass by reference because I try to code
// functionaly and passing by reference if a side efect
func (u *UserModel) InsertUser(user *User) error {
	existing, _ := u.find(user.Email)
	if existing.Id != "" {
		return errors.New("user exists")
	}
	user.Id = uuid.NewString()
	user.Password.Set()

	u.save(*user)
	return nil
}

// In nodejs I usually seperate the db code to a subdirectory in models
func (u *UserModel) find(email string) (User, error) {
	session := u.graphDb.NewReadSession()
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		parameters := map[string]interface{}{
			"email": email}

		records, err := tx.Run(
			"MATCH (u:User {email: $email}) RETURN u.first_name, u.last_name, u.email LIMIT 1",
			parameters)

		if err != nil {
			fmt.Printf("43: %#v", err)
			return nil, err
		}

		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		return User{
			FirstName: record.Values[0].(string),
			LastName:  record.Values[1].(string),
			Email:     record.Values[2].(string)}, nil

	})
	if err != nil {
		return User{}, err
	}
	return result.(User), nil
}

// In nodejs I usually seperate the db code to a subdirectory in models
func (u *UserModel) save(user User) error {
	session := u.graphDb.NewWriteSession()
	defer session.Close()

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		parameters := map[string]interface{}{
			"email":         user.Email,
			"first_name":    user.FirstName,
			"last_name":     user.LastName,
			"password_hash": bytes.NewBuffer(user.Password.Hash).String()}

		result, err := transaction.Run(
			"CREATE (:User {email: $email, first_name: $first_name, last_name: $last_name, password_hash: $password_hash})",
			parameters)

		if err != nil {
			fmt.Print(err.Error())
			return nil, err
		}

		return nil, result.Err()
	})
	return nil
}
