package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/guilhermerodrigues680/slack-bot-ripper-go/cmd/bot"

	"github.com/gorilla/mux"
)

func slackOutgoing(w http.ResponseWriter, r *http.Request) {

	outgoing := bot.SlackOutgoingMessage{
		UserName:    r.FormValue("user_name"),
		Text:        r.FormValue("text"),
		TriggerWord: r.FormValue("trigger_word"),
	}

	log.Printf("Usuario: %s, Comando: %s, Trigger: %s\n", outgoing.UserName, outgoing.Text, outgoing.TriggerWord)

	res := bot.RunCommand(outgoing)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Version é transmitido pelo ldflags durante a compilacao
var Version = "0.0.0-development"

// Build é transmitido pelo ldflags durante a compilacao
var Build = "0000000"

// BuildTime é transmitido pelo ldflags durante a compilacao
var BuildTime = "2009-11-10T23:00:00Z"

func main() {
	log.Printf("Version: %s\t(Build: %s)\tBuild Time: %s", Version, Build, BuildTime)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/slack/outgoing-webhook", slackOutgoing).Methods("POST")

	log.Println("Executando bot")
	http.ListenAndServe(":8080", r)
}
