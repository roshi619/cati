package command

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/roshi619/vbs"
)

// Configuration Precedence
// * viper.Set
// * flag
// * env
// * file
// * defaults

var baseDefaults = map[string]interface{}{
	"defaults": []string{"banner"},
	"message":  "Done!",

	"nsuser.soundName":     "Ping",
	"nsuser.soundNameFail": "Basso",

	"say.voice": "Alex",

	"espeak.voiceName": "english-us",

	"speechsynthesizer.voice": "Microsoft David Desktop",

	"bearychat.incomingHookURI": "",

	"hipchat.accessToken": "",
	"hipchat.room":        "",

	"pushbullet.accessToken": "",
	"pushbullet.deviceIden": "",

	"pushover.apiToken": "",
	"pushover.userKey":  "",

	"pushsafer.key": "",

	"simplepush.key":   "",
	"simplepush.event": "",

	"slack.token":    "",
	"slack.channel":  "",
	"slack.username": "cati",
}

func setCATIDefaults(v *viper.Viper) {
	for key, val := range baseDefaults {
		v.SetDefault(key, val)
	}
}

var keyEnvBindings = map[string]string{
	"nsuser.soundName":     "CATI_NSUSER_SOUNDNAME",
	"nsuser.soundNameFail": "CATI_NSUSER_SOUNDNAMEFAIL",

	"say.voice": "CATI_SAY_VOICE",

	"espeak.voiceName": "CATI_ESPEAK_VOICENAME",

	"speechsynthesizer.voice": "CATI_SPEECHSYNTHESIZER_VOICE",

	"bearychat.incomingHookURI": "CATI_BEARYCHAT_INCOMINGHOOKURI",

	"hipchat.accessToken": "CATI_HIPCHAT_ACCESSTOKEN",
	"hipchat.room":        "CATI_HIPCHAT_ROOM",

	"pushbullet.accessToken": "CATI_PUSHBULLET_ACCESSTOKEN",
	"pushbullet.deviceIden": "CATI_PUSHBULLET_DEVICEIDEN",

	"pushover.apiToken": "CATI_PUSHOVER_APITOKEN",
	"pushover.userKey":  "CATI_PUSHOVER_USERKEY",

	"pushsafer.key": "CATI_PUSHSAFER_KEY",

	"simplepush.key":   "CATI_SIMPLEPUSH_KEY",
	"simplepush.event": "CATI_SIMPLEPUSH_EVENT",

	"slack.token":    "CATI_SLACK_TOKEN",
	"slack.channel":  "CATI_SLACK_CHANNEL",
	"slack.username": "CATI_SLACK_USERNAME",
}

var keyEnvBindingsDeprecated = map[string]string{
	"CATI_NSUSER_SOUNDNAME":          "CATI_SOUND",
	"CATI_NSUSER_SOUNDNAMEFAIL":      "CATI_SOUND_FAIL",
	"CATI_SAY_VOICE":                 "CATI_VOICE",
	"CATI_ESPEAK_VOICENAME":          "CATI_VOICE",
	"CATI_SPEECHSYNTHESIZER_VOICE":   "CATI_VOICE",
	"CATI_BEARYCHAT_INCOMINGHOOKURI": "CATI_BC_INCOMING_URI",
	"CATI_HIPCHAT_ACCESSTOKEN":       "CATI_HIPCHAT_TOK",
	"CATI_HIPCHAT_ROOM":              "CATI_HIPCHAT_DEST",
	"CATI_PUSHBULLET_ACCESSTOKEN":    "CATI_PUSHBULLET_TOK",
	"CATI_PUSHOVER_TOKEN":            "CATI_PUSHOVER_TOK",
	"CATI_PUSHOVER_USER":             "CATI_PUSHOVER_DEST",
	"CATI_SLACK_TOKEN":               "CATI_SLACK_TOK",
	"CATI_SLACK_CHANNEL":             "CATI_SLACK_DEST",
}

func bindCATIEnv(v *viper.Viper) error {
	for key, val := range keyEnvBindings {
		if err := v.BindEnv(key, val); err != nil {
			return err
		}
	}

	// Map old deprecated env vars to new ones.
	for newEnv, oldEnv := range keyEnvBindingsDeprecated {
		v := os.Getenv(oldEnv)
		if v == "" {
			continue
		}

		fmt.Fprintf(os.Stderr, "Warning: %s is deprecated, use %s instead\n",
			oldEnv, newEnv)
		fmt.Fprintf(os.Stderr, "Remapping %s=%s to %s\n", oldEnv, v, newEnv)

		if err := os.Setenv(newEnv, v); err != nil {
			return err
		}
	}

	return nil
}

func setupConfigFile(fileFlag string, v *viper.Viper) error {
	viper.SupportedExts = []string{"yaml"}
	var configPaths []string

	if fileFlag != "" {
		configPaths = append(configPaths, fileFlag)
	}

	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfig == "" {
		xdgConfig = filepath.Join(os.ExpandEnv("$HOME"), ".config", "cati", "cati.yaml")
	} else {
		xdgConfig = filepath.Join(xdgConfig, "cati", "cati.yaml")
	}

	configPaths = append(configPaths,
		filepath.Join(".", ".cati.yaml"),
		xdgConfig,
	)

	var config io.Reader
	var errMsg []string
	for _, p := range configPaths {
		data, err := ioutil.ReadFile(p)
		if err != nil {
			errMsg = append(errMsg, err.Error())
			continue
		}

		config = bytes.NewReader(data)
		break
	}
	if config == nil {
		return fmt.Errorf("failed to read config file: %v", errMsg)
	}

	v.SetConfigType("yaml")
	return v.ReadConfig(config)
}

// configureApp merges together different configuration sources.
func configureApp(v *viper.Viper, flags *pflag.FlagSet) error {
	setCATIDefaults(v)

	if err := bindCATIEnv(v); err != nil {
		return err
	}

	// Don't care about this error, fileFlag can be blank.
	fileFlag, _ := flags.GetString("file")
	if err := setupConfigFile(fileFlag, v); err != nil {
		// Not the end of the world if we can't read the config file.
		vbs.Println(err)
	}

	if flags == nil {
		return nil
	}

	return v.BindPFlag("message", flags.Lookup("message"))
}

func enabledFromSlice(defaults []string) map[string]bool {
	// defaults should come from viper, which should  have processed baseDefaults
	// and config file values.

	services := map[string]bool{
		"banner":     false,
		"bearychat":  false,
		"hipchat":    false,
		"pushbullet": false,
		"pushover":   false,
		"pushsafer":  false,
		"simplepush": false,
		"slack":      false,
		"speech":     false,
	}

	for _, name := range defaults {
		// Check if name is in services to avoid bad names from getting added
		// to map.
		if _, ok := services[name]; ok {
			services[name] = true
		}
	}

	return services
}

func hasServiceFlags(flags *pflag.FlagSet) bool {
	services := map[string]bool{
		"banner":     false,
		"bearychat":  false,
		"hipchat":    false,
		"pushbullet": false,
		"pushover":   false,
		"pushsafer":  false,
		"simplepush": false,
		"slack":      false,
		"speech":     false,
	}

	flags.Visit(func(f *pflag.Flag) {
		if _, ok := services[f.Name]; ok {
			services[f.Name] = true
		}
	})

	for _, enabled := range services {
		if enabled {
			return true
		}
	}
	return false
}

func enabledFromFlags(flags *pflag.FlagSet) map[string]bool {
	services := map[string]bool{
		"banner":     false,
		"bearychat":  false,
		"hipchat":    false,
		"pushbullet": false,
		"pushover":   false,
		"pushsafer":  false,
		"simplepush": false,
		"slack":      false,
		"speech":     false,
	}

	// Visit flags that have been set.
	flags.Visit(func(f *pflag.Flag) {
		// pflag normalizes false, f, 0 to "false".
		if f.Value.Type() == "bool" && f.Value.String() == "false" {
			// Skip bool flags that are set to false.
			return
		}

		// Ignore flags that aren't service names.
		if _, ok := services[f.Name]; ok {
			services[f.Name] = true
		}
	})

	return services
}

func enabledServices(v *viper.Viper, flags *pflag.FlagSet) map[string]struct{} {
	var services map[string]bool

	if hasServiceFlags(flags) {
		// Highest precedence.
		services = enabledFromFlags(flags)
	} else if s := os.Getenv("CATI_DEFAULT"); s != "" {
		services = enabledFromSlice(strings.Split(s, " "))
	} else if s := v.GetStringSlice("defaults"); len(s) != 0 {
		// Lowest precedence.
		services = enabledFromSlice(s)
	}

	filtered := make(map[string]struct{})
	for service, enabled := range services {
		if enabled {
			filtered[service] = struct{}{}
		}
	}

	return filtered
}

func getCATIfications(v *viper.Viper, services map[string]struct{}) []catification {
	title := v.GetString("title")
	message := v.GetString("message")

	var catis []catification

	if _, ok := services["banner"]; ok {
		catis = append(catis, getBanner(title, message, v))
	}

	if _, ok := services["speech"]; ok {
		catis = append(catis, getSpeech(title, message, v))
	}

	if _, ok := services["bearychat"]; ok {
		catis = append(catis, getBearyChat(title, message, v))
	}

	if _, ok := services["hipchat"]; ok {
		catis = append(catis, getHipChat(title, message, v))
	}

	if _, ok := services["pushbullet"]; ok {
		catis = append(catis, getPushbullet(title, message, v))
	}

	if _, ok := services["pushover"]; ok {
		catis = append(catis, getPushover(title, message, v))
	}

	if _, ok := services["pushsafer"]; ok {
		catis = append(catis, getPushsafer(title, message, v))
	}

	if _, ok := services["simplepush"]; ok {
		catis = append(catis, getSimplepush(title, message, v))
	}

	if _, ok := services["slack"]; ok {
		catis = append(catis, getSlack(title, message, v))
	}

	return catis
}
