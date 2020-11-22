package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type slackOutgoingMessage struct {
	UserName    string `json:"username"`
	Text        string `json:"text"`
	TriggerWord string `json:"triggerword"`
}

type slackOutgoingResponse struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

func slackOutgoing(w http.ResponseWriter, r *http.Request) {

	outgoing := slackOutgoingMessage{
		UserName:    r.FormValue("user_name"),
		Text:        r.FormValue("text"),
		TriggerWord: r.FormValue("trigger_word"),
	}

	fmt.Printf("Usuario: %s, Comando: %s, Trigger: %s\n", outgoing.UserName, outgoing.Text, outgoing.TriggerWord)

	res := slackOutgoingResponse{
		Text:     "Go Vive",
		Username: "Golang BOT",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/slack/outgoing-webhook", slackOutgoing).Methods("POST")

	fmt.Println("Executando bot")
	http.ListenAndServe(":8080", r)
}
