package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type yesNoReponse struct {
	Answer string `json:answer`
	Forced bool   `json:forced`
	Image  string `json:image`
}

func getTextWithoutTriggerWord(text string, triggerWord string) string {
	regexTriggerWord, _ := regexp.Compile(triggerWord + "\\s?")
	command := regexTriggerWord.ReplaceAllString(text, "")

	return command
}

func getFirstNameFromUserName(userName string) string {
	firstName := strings.Split(userName, ".")[0]
	return strings.Title(strings.ToLower(firstName))
}

func commandYesNo(command string, userName string) (SlackOutgoingResponse, error) {
	resp, err := http.Get("https://yesno.wtf/api") // ?force=maybe
	if err != nil {
		log.Println("Erro ao buscar api")
		return SlackOutgoingResponse{Text: "", Username: ""}, err
	}

	var result yesNoReponse
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)

	if result.Answer == "yes" {
		return SlackOutgoingResponse{command[1:] + "\nestou pensando... <" + result.Image + "|Sim!>", "Golang BOT"}, nil
	}

	return SlackOutgoingResponse{command[1:] + "\nestou pensando... <" + result.Image + "|Não!>", "Golang BOT"}, nil
}

func executeCommand(command string, userName string) SlackOutgoingResponse {

	if strings.HasPrefix(command, "?") && strings.HasSuffix(command, "?") {
		resp, _ := commandYesNo(command, userName)
		return resp
	}

	switch strings.ToUpper(command) {
	case "TA AI?", "TA VIVO?":
		return SlackOutgoingResponse{"Tô aqui de pé " + userName + ", haha!", "Golang BOT"}
	case "PARTIU", "BORA":
		return SlackOutgoingResponse{"Já tá na hora? Uhuuu, então partiu!", "Golang BOT"}
	default:
		return SlackOutgoingResponse{"Poxa " + userName + ", não conheco o comando `" + command + "` :(", "Golang BOT"}
	}
}

// RunCommand processa um comando enviado ao bot
func RunCommand(outgoing SlackOutgoingMessage) SlackOutgoingResponse {
	command := getTextWithoutTriggerWord(outgoing.Text, outgoing.TriggerWord)
	log.Printf("Comando recebido: %s", command)
	firstName := getFirstNameFromUserName(outgoing.UserName)
	return executeCommand(command, firstName)
}
