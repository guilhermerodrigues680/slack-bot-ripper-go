package bot

import (
	"log"
	"regexp"
	"strings"
)

func getTextWithoutTriggerWord(text string, triggerWord string) string {
	regexTriggerWord, _ := regexp.Compile(triggerWord + "\\s?")
	command := regexTriggerWord.ReplaceAllString(text, "")

	return command
}

func executeCommand(command string, userName string) SlackOutgoingResponse {
	switch strings.ToUpper(command) {
	case "TA AI?":
		fallthrough
	case "TA VIVO?":
		return SlackOutgoingResponse{
			Text:     "To aqui de pé haha!",
			Username: "Golang BOT",
		}
	case "PARTIU":
		fallthrough
	case "BORA":
		return SlackOutgoingResponse{
			Text:     "Já tá na hora? Uhuuu, então partiu!",
			Username: "Golang BOT",
		}
	default:
		return SlackOutgoingResponse{
			Text:     "Poxa " + userName + ", não conheco o comando `" + command + "` :(",
			Username: "Golang BOT",
		}
	}
}

// RunCommand processa um comando enviado ao bot
func RunCommand(outgoing SlackOutgoingMessage) SlackOutgoingResponse {
	command := getTextWithoutTriggerWord(outgoing.Text, outgoing.TriggerWord)
	log.Printf("Comando recebido: %s", command)
	return executeCommand(command, outgoing.UserName)
}
