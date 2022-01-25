package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewHealthCheckRoute() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`"status":"ok"`))
	})
}
