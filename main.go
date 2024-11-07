package main

import (
	"api-cotacoes/api"
	"api-cotacoes/handlers"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	tmpl, _ := template.ParseFiles("web/templates/index.html")

	var rotas handlers.StructRotas = handlers.StructRotas{
		Rotas:    mux.NewRouter(),
		Template: tmpl,
	}

	var api api.Api = api.Api{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	rotas.InitAll(api)
	rotas.Run()
}
