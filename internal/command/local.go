// +build !darwin
// +build !windows

package command

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/roshi619/cati/service/espeak"
	"github.com/roshi619/cati/service/freedesktop"
)

func getBanner(title, message string, v *viper.Viper) catification {
	return &freedesktop.CATIfication{
		Summary:       title,
		Body:          message,
		ExpireTimeout: 500,
		AppIcon:       "utilities-terminal",
	}
}

func getSpeech(title, message string, v *viper.Viper) catification {
	return &espeak.CATIfication{
		Text:      fmt.Sprintf("%s %s", title, message),
		VoiceName: v.GetString("espeak.voiceName"),
	}
}
