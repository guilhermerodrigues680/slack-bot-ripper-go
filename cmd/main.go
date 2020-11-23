package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/guilhermerodrigues680/slack-bot-ripper-go/cmd/bot"
	"github.com/guilhermerodrigues680/slack-bot-ripper-go/cmd/stopwatch"

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

func meusTestes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()["q"][0]
	fmt.Println(q)

	switch q {
	case "start":
		stopwatchEl.Start()
	case "pause":
		stopwatchEl.Pause()
	case "reset":
		stopwatchEl.Reset()
	case "status":
		json.NewEncoder(w).Encode(stopwatchEl.Status())
		return
	}

	fmt.Fprintf(w, "Comando: %s", q)
}

// Version é transmitido pelo ldflags durante a compilacao
var Version = "0.0.0-development"

// Build é transmitido pelo ldflags durante a compilacao
var Build = "0000000"

// BuildTime é transmitido pelo ldflags durante a compilacao
var BuildTime = "2009-11-10T23:00:00Z"

var stopwatchEl = stopwatch.NewStopwatch()

func main() {
	log.Printf("Version: %s\t(Build: %s)\tBuild Time: %s", Version, Build, BuildTime)

	go stopwatchEl.Cronometro() // Go routine que vive junto com o processo

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/slack/outgoing-webhook", slackOutgoing).Methods("POST")
	r.HandleFunc("/teste", meusTestes).Methods("GET")

	log.Println("Executando bot")
	http.ListenAndServe(":8080", r)
}
