package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"html/template"

	"api-cotacoes/api"
	"log"

	"github.com/gorilla/mux"
	// "net/http"
	// "encoding/json"
)

type StructRotas struct {
	Rotas    *mux.Router
	Template *template.Template
}

type Moedas struct {
	AOA float64 `json:"AOA"`
	BRL float64 `json:"BRL"`
	EUR float64 `json:"EUR"`
}

// variaveis que serao acessados pelas funcoes abaixo
var data_moedas Moedas

func (r *StructRotas) InitAll(api_client api.Api) {
	// setar static files
	r.Rotas.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	// setar pagina inicial
	r.Rotas.HandleFunc("/", func(w http.ResponseWriter, client *http.Request) {
		r.Template.Execute(w, nil)
	})

	// setar rota da api
	r.Rotas.HandleFunc("/api", func(w http.ResponseWriter, client *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		status, data_full := api_client.Rates()

		if status {
			// se true, converte para ResponseXe
			var data_f api.ResponseXe = data_full.(api.ResponseXe)

			// pega apenas os dados necessarios
			data_moedas = Moedas{
				AOA: data_f.Rates["AOA"],
				BRL: data_f.Rates["BRL"],
				EUR: data_f.Rates["EUR"],
			}

			// enviar o JSON
			json.NewEncoder(w).Encode(data_moedas)
		} else {
			// se false, envie o map direto
			json.NewEncoder(w).Encode(data_full)
		}
	})
}

func (r *StructRotas) Run() {
	fmt.Println("Webservice iniciado!")
	err := http.ListenAndServe(":80", r.Rotas)
	if err != nil {
		log.Fatal(err)
	}
}
