package models

type User struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Password  Password `json:"-"`
}

// With node I would validate with a JSON schema validator
func (u User) Validate() []string {
	return nil
}
