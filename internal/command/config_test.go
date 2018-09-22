package command

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func countSettingsKeys(t *testing.T, m map[string]interface{}) int {
	t.Helper()

	var keys int
	for _, v := range m {
		if sub, ok := v.(map[string]interface{}); ok {
			// Don't count the object, just its keys.
			keys += len(sub)
		}

		if _, ok := v.(string); ok {
			// v is just a string key.
			keys++
		}

		if _, ok := v.([]string); ok {
			// v is just a string key.
			keys++
		}
	}
	return keys
}

func TestSetCATIDefaults(t *testing.T) {
	v := viper.New()
	setCATIDefaults(v)

	haveKeys := countSettingsKeys(t, v.AllSettings())
	if haveKeys != len(baseDefaults) {
		t.Error("Unexpected base config length")
		t.Errorf("have=%d; want=%d", haveKeys, len(baseDefaults))
	}
}

func getCATIEnv(t *testing.T) map[string]string {
	t.Helper()

	catiEnv := make(map[string]string)
	for _, env := range keyEnvBindings {
		catiEnv[env] = os.Getenv(env)
	}
	return catiEnv
}

func clearCATIEnv(t *testing.T) {
	t.Helper()

	for _, env := range keyEnvBindings {
		if err := os.Unsetenv(env); err != nil {
			t.Fatalf("failed to clear cati env: %s", err)
		}
	}
}

func setCATIEnv(t *testing.T, m map[string]string) {
	t.Helper()

	for env, val := range m {
		if err := os.Setenv(env, val); err != nil {
			t.Fatalf("failed to set cati env: %s", err)
		}
	}
}

func TestBindCATIEnv(t *testing.T) {
	orig := getCATIEnv(t)
	defer setCATIEnv(t, orig)

	clearCATIEnv(t)

	v := viper.New()
	bindCATIEnv(v)

	haveKeys := countSettingsKeys(t, v.AllSettings())
	if haveKeys != 0 {
		t.Error("Environment should be cleared")
		t.Error(v.AllSettings())
	}

	var numSet int
	for _, env := range keyEnvBindings {
		if err := os.Setenv(env, "foo"); err != nil {
			t.Errorf("Setenv error: %s", err)
			continue
		}
		numSet++
	}

	haveKeys = countSettingsKeys(t, v.AllSettings())
	wantKeys := numSet
	if haveKeys != wantKeys {
		t.Error("Unexpected base config length")
		t.Errorf("have=%d; want=%d", haveKeys, wantKeys)
		t.Error(v.AllSettings())
	}
}

func TestSetupConfigFile(t *testing.T) {
	v := viper.New()
	if err := setupConfigFile("testdata/cati.yaml", v); err != nil {
		t.Error(err)
	}

	const want = 1
	have := countSettingsKeys(t, v.AllSettings())
	if have != want {
		t.Error("Unexpected number of keys")
		t.Errorf("have=%d; want=%d", have, want)
	}
}

func TestConfigureApp(t *testing.T) {
	orig := getCATIEnv(t)
	defer setCATIEnv(t, orig)

	cases := []struct {
		name       string
		configFile string
		env        string
		want       string
	}{
		{
			// Config file should take precedence.
			name:       "defaults and file",
			configFile: "testdata/cati.yaml",
			want:       "testSoundName",
		},
		{
			// Env should take precedence.
			name:       "defaults, file, and env",
			configFile: "testdata/cati.yaml",
			env:        "CATI_NSUSER_SOUNDNAME",
			want:       "testSoundName",
		},
		{
			// Defaults should take precedence.
			name: "defaults",
			want: baseDefaults["nsuser.soundName"].(string),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			clearCATIEnv(t)

			v := viper.New()
			flags := pflag.NewFlagSet("testconfigureapp", pflag.ContinueOnError)
			defineFlags(flags)

			if c.configFile != "" {
				flags.Set("file", c.configFile)
			}
			if c.env != "" {
				if err := os.Setenv(c.env, c.want); err != nil {
					t.Errorf("Failed to set env: %s", err)
				}
			}

			if err := configureApp(v, flags); err != nil {
				t.Error(err)
			}

			have := v.GetString("nsuser.soundName")
			if have != c.want {
				t.Error("Unexpected config value")
				t.Errorf("have=%s; want=%s", have, c.want)
				t.Error("nsuser:", v.Sub("nsuser").AllSettings())
			}
		})
	}
}

func TestEnabledServices(t *testing.T) {
	orig := getCATIEnv(t)
	defer setCATIEnv(t, orig)
	clearCATIEnv(t)

	t.Run("flag override", func(t *testing.T) {
		v := viper.New()
		// For tests, we prepend the testdata dir so that we check for a config
		// file there first.
		v.AddConfigPath("testdata")

		flags := pflag.NewFlagSet("testenabledservices", pflag.ContinueOnError)
		defineFlags(flags)

		configureApp(v, flags)

		want := true
		flags.Set("slack", fmt.Sprint(want))
		services := enabledServices(v, flags)

		if len(services) != 1 {
			t.Error("Unexpected number of enabled services")
			t.Errorf("have=%d; want=%d", len(services), 1)
		}

		_, have := services["slack"]
		if have != want {
			t.Error("Unexpected enabled state")
			t.Errorf("have=%t; want=%t", have, want)
		}
	})

	t.Run("non-service flags", func(t *testing.T) {
		v := viper.New()
		// For tests, we prepend the testdata dir so that we check for a config
		// file there first.
		v.AddConfigPath("testdata")

		flags := pflag.NewFlagSet("testenabledservices", pflag.ContinueOnError)
		defineFlags(flags)

		configureApp(v, flags)

		flags.Set("verbose", "true")
		services := enabledServices(v, flags)

		// We should end up taking the defaults.

		if len(services) != 1 {
			t.Error("Unexpected number of enabled services")
			t.Errorf("have=%d; want=%d", len(services), 1)
			t.Error("services:", services)
		}

		want := true
		_, have := services["banner"]
		if have != want {
			t.Error("Unexpected enabled state")
			t.Errorf("have=%t; want=%t", have, want)
		}
	})

	t.Run("env override", func(t *testing.T) {
		v := viper.New()
		// For tests, we prepend the testdata dir so that we check for a config
		// file there first.
		v.AddConfigPath("testdata")

		flags := pflag.NewFlagSet("testenabledservices", pflag.ContinueOnError)
		defineFlags(flags)

		configureApp(v, flags)

		if err := os.Setenv("CATI_DEFAULT", "slack"); err != nil {
			t.Fatal(err)
		}
		defer os.Unsetenv("CATI_DEFAULT")

		services := enabledServices(v, flags)

		if len(services) != 1 {
			t.Error("Unexpected number of enabled services")
			t.Errorf("have=%d; want=%d", len(services), 1)
		}

		_, have := services["slack"]
		want := true
		if have != want {
			t.Error("Unexpected enabled state")
			t.Errorf("have=%t; want=%t", have, want)
		}
	})

	t.Run("defaults", func(t *testing.T) {
		v := viper.New()
		// For tests, we prepend the testdata dir so that we check for a config
		// file there first.
		v.AddConfigPath("testdata")

		flags := pflag.NewFlagSet("testenabledservices", pflag.ContinueOnError)
		defineFlags(flags)

		configureApp(v, flags)

		services := enabledServices(v, flags)

		if len(services) != 1 {
			t.Error("Unexpected number of enabled services")
			t.Errorf("have=%d; want=%d", len(services), 1)
		}

		_, have := services["banner"]
		want := true
		if have != want {
			t.Error("Unexpected enabled state")
			t.Errorf("have=%t; want=%t", have, want)
		}
	})
}

func TestGetCATIfications(t *testing.T) {
	services := []string{
		"banner",
		"bearychat",
		"hipchat",
		"pushbullet",
		"pushover",
		"pushsafer",
		"simplepush",
		"slack",
		"speech",
	}

	for _, name := range services {
		t.Run(fmt.Sprintf("get %s catification", name), func(t *testing.T) {
			v := viper.New()
			s := map[string]struct{}{name: struct{}{}}

			catis := getCATIfications(v, s)
			if len(catis) != 1 {
				t.Error("Unexpected number of catifications")
				t.Errorf("have=%d; want=%d", len(catis), 1)
			}
		})
	}
}
