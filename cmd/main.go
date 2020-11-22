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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/slack/outgoing-webhook", slackOutgoing).Methods("POST")

	log.Println("Executando bot")
	http.ListenAndServe(":8080", r)
}
