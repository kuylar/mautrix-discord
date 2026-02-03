package plural

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPluRalMessageInfo(apiKey string, channelId string, messageId string) (*PluRalMessage, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.plural.gg/messages/%s/%s?member=true", channelId, messageId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("User-Agent", "kuylar_mautrix-discord (mailto:kuylar@kuylar.dev, https://github.com/kuylar/mautrix-discord)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// idk
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	var pluralMessage PluRalMessage
	err = json.NewDecoder(res.Body).Decode(&pluralMessage)
	if err != nil {
		return nil, err
	}
	return &pluralMessage, nil
}

type PluRalMessage struct {
	OriginalID  string       `json:"original_id"`
	ProxyID     string       `json:"proxy_id"`
	AuthorID    string       `json:"author_id"`
	ChannelID   string       `json:"channel_id"`
	MemberID    string       `json:"member_id"`
	Reason      string       `json:"reason"`
	WebhookID   string       `json:"webhook_id"`
	ReferenceID string       `json:"reference_id"`
	Member      PluRalMember `json:"member"`
}

type PluRalMember struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Pronouns  string `json:"pronouns"`
	Bio       string `json:"bio"`
	Birthday  string `json:"birthday"`
	Color     int64  `json:"color"`
	AvatarURL string `json:"avatar_url"`
	Supporter bool   `json:"supporter"`
	Private   bool   `json:"private"`
}
