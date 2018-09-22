package command

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/roshi619/cati/service/bearychat"
	"github.com/roshi619/cati/service/hipchat"
	"github.com/roshi619/cati/service/pushbullet"
	"github.com/roshi619/cati/service/pushover"
	"github.com/roshi619/cati/service/pushsafer"
	"github.com/roshi619/cati/service/simplepush"
	"github.com/roshi619/cati/service/slack"
)

var httpClient = &http.Client{Timeout: 30 * time.Second}

func getBearyChat(title, message string, v *viper.Viper) catification {
	return &bearychat.CATIfication{
		Text:            fmt.Sprintf("**%s**\n%s", title, message),
		IncomingHookURI: v.GetString("bearychat.incomingHookURI"),
		Client:          httpClient,
	}
}

func getHipChat(title, message string, v *viper.Viper) catification {
	return &hipchat.CATIfication{
		AccessToken:   v.GetString("hipchat.accessToken"),
		Room:          v.GetString("hipchat.room"),
		Message:       fmt.Sprintf("%s\n%s", title, message),
		MessageFormat: "text",
		Client:        httpClient,
	}
}

func getPushbullet(title, message string, v *viper.Viper) catification {
	return &pushbullet.CATIfication{
		Title:       title,
		Body:        message,
		Type:        "note",
		AccessToken: v.GetString("pushbullet.accessToken"),
		DeviceIden:  v.GetString("pushbullet.deviceIden"),
		Client:      httpClient,
	}
}

func getPushover(title, message string, v *viper.Viper) catification {
	return &pushover.CATIfication{
		Title:    title,
		Message:  message,
		APIToken: v.GetString("pushover.apiToken"),
		UserKey:  v.GetString("pushover.userKey"),
		Client:   httpClient,
	}
}

func getPushsafer(title, message string, v *viper.Viper) catification {
	return &pushsafer.CATIfication{
		Title:   title,
		Message: message,
		Key:     v.GetString("pushsafer.key"),
		Client:  httpClient,
	}
}

func getSimplepush(title, message string, v *viper.Viper) catification {
	return &simplepush.CATIfication{
		Title:   title,
		Message: message,
		Key:     v.GetString("simplepush.key"),
		Event:   v.GetString("simplepush.event"),
		Client:  httpClient,
	}
}

func getSlack(title, message string, v *viper.Viper) catification {
	return &slack.CATIfication{
		Token:     v.GetString("slack.token"),
		Channel:   v.GetString("slack.channel"),
		Username:  v.GetString("slack.username"),
		Text:      fmt.Sprintf("%s\n%s", title, message),
		IconEmoji: ":rocket:",

		Client: httpClient,
	}
}
