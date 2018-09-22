package command

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/roshi619/cati/service/nsuser"
	"github.com/roshi619/cati/service/say"
)

func getBanner(title, message string, v *viper.Viper) catification {
	return &nsuser.CATIfication{
		Title:           title,
		InformativeText: message,
		SoundName:       v.GetString("nsuser.soundName"),
	}
}

func getSpeech(title, message string, v *viper.Viper) catification {
	return &say.CATIfication{
		Voice: v.GetString("say.voice"),
		Text:  fmt.Sprintf("%s %s", title, message),
		Rate:  200,
	}
}
