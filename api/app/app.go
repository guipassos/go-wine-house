package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guipassos/go-wine-house/api/app/handler"
	"github.com/guipassos/go-wine-house/api/app/model"
	"github.com/guipassos/go-wine-house/api/config"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	a.Get("/wines", a.GetAllWines)
	a.Post("/wines", a.CreateWine)
	a.Get("/wines/{id}", a.GetWine)
	a.Put("/wines/{id}", a.UpdateWine)
	a.Delete("/wines/{id}", a.DeleteWine)
	a.Put("/wines/{id}/disable", a.DisableWine)
	a.Put("/wines/{id}/enable", a.EnableWine)
}

// Wrap the router for GET method
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

// Handlers to manage Wine Data
func (a *App) GetAllWines(w http.ResponseWriter, r *http.Request) {
	handler.GetAllWines(a.DB, w, r)
}

func (a *App) CreateWine(w http.ResponseWriter, r *http.Request) {
	handler.CreateWine(a.DB, w, r)
}

func (a *App) GetWine(w http.ResponseWriter, r *http.Request) {
	handler.GetWine(a.DB, w, r)
}

func (a *App) UpdateWine(w http.ResponseWriter, r *http.Request) {
	handler.UpdateWine(a.DB, w, r)
}

func (a *App) DeleteWine(w http.ResponseWriter, r *http.Request) {
	handler.DeleteWine(a.DB, w, r)
}

func (a *App) DisableWine(w http.ResponseWriter, r *http.Request) {
	handler.DisableWine(a.DB, w, r)
}

func (a *App) EnableWine(w http.ResponseWriter, r *http.Request) {
	handler.EnableWine(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
