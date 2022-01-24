package routes

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewSignupRoute(userModel *models.UserModel, tokenModel *models.TokenModel) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//with node express i would use validate with
		var input struct {
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Password  string `json:"password`
		}

		// With node express I would use the body parser middlewhare
		json.NewDecoder(r.Body).Decode(&input)
		user := models.User{
			Email:     input.Email,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Password: models.Password{
				PlainText: input.Password}}

		errors := userModel.InsertUser(&user)
		if errors != nil {
			//TODO: add error response helper
			fmt.Fprintf(w, "%+v\n", errors)
			return
		}

		token, errors := tokenModel.GenerateJwtToken(user)
		if errors != nil {
			fmt.Fprintf(w, "%+v\n", errors)
			return
		}

		//TODO: add response helper
		w.Header().Set("Content-Type", "application/json")
		payload := fmt.Sprintf(`{"token":"%s"}`, token)
		w.Write([]byte(payload))
	})
}
