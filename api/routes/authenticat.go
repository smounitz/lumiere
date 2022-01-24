package routes

import (
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewLoginRoute(userModel *models.UserModel) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// With nodejs I would validate with a JSON schema validator
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password`
		}

		// With node express I would use the body parser middlewhare
		json.NewDecoder(r.Body).Decode(&input)
		//TODO:
	})
}
