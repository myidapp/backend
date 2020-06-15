package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	h "github.com/myidapp/backend/handlers"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	db, err := gorm.Open("sqlite3", "myid.db")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = db
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) Run(host string) {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"http://localhost:8001"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowCredentials(),
	)(a.Router)

	log.Fatal(http.ListenAndServe(host, cors))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) setRouters() {
	a.Post("/api/clear", a.Clear)
	a.Post("/api/sign", a.PostSign)
	a.Get("/api/sign/{uuid}", a.GetSign)
	a.Post("/api/schema", a.PostSchema)
	a.Get("/api/schema/{uuid}", a.GetSchema)
}

func (a *App) Clear(w http.ResponseWriter, r *http.Request) {
	h.Clear(a.DB, w, r)
}

func (a *App) GetSign(w http.ResponseWriter, r *http.Request) {
	h.GetSign(a.DB, w, r)
}

func (a *App) PostSign(w http.ResponseWriter, r *http.Request) {
	h.PostSign(a.DB, w, r)
}

func (a *App) GetSchema(w http.ResponseWriter, r *http.Request) {
	h.GetSchema(a.DB, w, r)
}

func (a *App) PostSchema(w http.ResponseWriter, r *http.Request) {
	h.PostSchema(a.DB, w, r)
}
