package structures

type (
	MessageSendRequest struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}

	MessageSendResponse struct {
		OK bool `json:"ok"`
		Result struct {
			MessageID int `json:"message_id`
			From struct {
				ID int `json:"id"`
				IsBot bool `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username string `json:"username"`
			}
			Chat struct {
				ID int `json:"id"`
				FirstName string `json:"first_name"`
				LastName string `json:"last_name"`
				Username string `json:"username"`
				Type string `json:"type"`
			}
			Date int `json:"date"`
			Text string `json:"text"`
		}
	}
)