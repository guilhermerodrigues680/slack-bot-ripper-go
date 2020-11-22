package bot

// SlackOutgoingMessage representa a mensagem enviada ao bot
type SlackOutgoingMessage struct {
	UserName    string `json:"username"`
	Text        string `json:"text"`
	TriggerWord string `json:"triggerword"`
}

// SlackOutgoingResponse representa a mensagem de resposta do bot
type SlackOutgoingResponse struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}
