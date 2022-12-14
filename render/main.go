package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "snack", Email: "snack@naver.com"}

	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "snack", Email: "snack@naver.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory:  "render/template",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	//New(NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("render/public")))
	n.UseHandler(mux)

	http.ListenAndServe(":3338", n)
}
