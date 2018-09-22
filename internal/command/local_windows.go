package command

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/roshi619/cati/service/catifyicon"
	"github.com/roshi619/cati/service/speechsynthesizer"
)

func getBanner(title, message string, v *viper.Viper) catification {
	return &catifyicon.CATIfication{
		BalloonTipTitle: title,
		BalloonTipText:  message,
		BalloonTipIcon:  catifyicon.BalloonTipIconInfo,
	}
}

func getSpeech(title, message string, v *viper.Viper) catification {
	return &speechsynthesizer.CATIfication{
		Text:  fmt.Sprintf("%s %s", title, message),
		Rate:  3,
		Voice: v.GetString("speechsynthesizer.voice"),
	}
}
