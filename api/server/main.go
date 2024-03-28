package main

import (
	"Multithreading/api/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

import "github.com/go-chi/chi"

func main() {
	r := chi.NewRouter()

	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", GetCep)
	})

	http.ListenAndServe(":8000", r)
}

func GetCep(w http.ResponseWriter, r *http.Request) {
	log.Println("Request started")
	defer log.Println("Request ended")

	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c1 := make(chan dto.Message)
	c2 := make(chan dto.Message)

	go callBrasilAPI(cep, c1)
	go callViaCEPAPI(cep, c2)

	var brasilAPIResponse, viaCEPResponse dto.Message

	select {
	case brasilAPIResponse = <-c1:
		if brasilAPIResponse.API == "BrasilAPI" {
			if brasilAPIResponse.Msg != "" {
				json.NewEncoder(w).Encode(brasilAPIResponse)
				return
			}
		}
	case viaCEPResponse = <-c2:
		if viaCEPResponse.API == "ViaCEP" {
			if viaCEPResponse.Msg != "" {
				json.NewEncoder(w).Encode(viaCEPResponse)
				return
			}
		}
	case <-time.After(time.Second):
		fmt.Fprint(w, "Timeout")
		return
	}
	http.Error(w, "CEP não encontrado ou erro na consulta", http.StatusNotFound)
}

func callBrasilAPI(cep string, c chan dto.Message) {
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		c <- dto.Message{"BrasilAPI", err.Error()}
		return
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		c <- dto.Message{"BrasilAPI", fmt.Sprintf("Erro na consulta: %s", req.Status)}
		return
	}

	resBrasilAPI, err := io.ReadAll(req.Body)
	if err != nil {
		c <- dto.Message{"BrasilAPI", err.Error()}
		return
	}

	c <- dto.Message{"BrasilAPI", string(resBrasilAPI)}
}

func callViaCEPAPI(cep string, c chan dto.Message) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		c <- dto.Message{"ViaCEP", err.Error()}
		return
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		c <- dto.Message{"ViaCEP", fmt.Sprintf("Erro na consulta: %s", req.Status)}
		return
	}

	resViaCEPAPI, err := io.ReadAll(req.Body)
	if err != nil {
		c <- dto.Message{"ViaCEP", err.Error()}
		return
	}

	c <- dto.Message{"ViaCEP", string(resViaCEPAPI)}
}
