% CATI.YAML(5) cati 3.1.0 | CATI Configuration File Format
% Jaime Piña
% 2018/03/25

#  NAME

cati.yaml - cati configuration file

# SYNOPSIS

cati.yaml

# DESCRIPTION

File format is YAML.

If not explicitly set with \--file, then cati will check the following paths,
in the following order.

* ./.cati.yaml
* $XDG_CONFIG_HOME/cati/cati.yaml

If $XDG_CONFIG_HOME is empty, then $HOME/.config will be used as its default
value and cati will check $HOME/.config/cati/cati.yaml.

# NSUSER

soundName
: Banner success sound. Default is Ping. Possible options are Basso, Blow,
  Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
  Submarine, Tink. See /System/Library/Sounds for available sounds.

soundNameFail
: Banner failure sound. Default is Basso. Possible options are Basso,
  Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
  Submarine, Tink. See /System/Library/Sounds for available sounds.

# SAY

voice
: Name of voice used for speech catifications.

# ESPEAK

voiceName
: Name of voice used for speech catifications.

# SPEECHSYNTHESIZER

voice
: Name of voice used for speech catifications.

# BEARYCHAT

incomingHookURI
: BearyChat incoming URI.

# HIPCHAT

accessToken
: HipChat access token. Log into your HipChat account and retrieve a token
  from the Room CATIfication Tokens page.

room
: HipChat message destination. Can be either a Room name or ID.

# PUSHBULLET

accessToken
: Pushbullet access token. Log into your Pushbullet account and retrieve a
  token from the Account Settings page.

deviceIden
: Pushbullet device iden of the target device, if sending to a single device.

# PUSHOVER

apiToken
: Pushover access token. Log into your Pushover account and create a
  token from the Create New Application/Plugin page.

userKey
: Pushover message destination. Should be your User Key.

# PUSHSAFER

key
: Pushsafer private or alias key. Log into your Pushsafer account and note
  your private or alias key.

# SIMPLEPUSH

key
: Simplepush key. Install the Simplepush app and retrieve your key there.

event
: Customize ringtone and vibration.

# SLACK

token
: Slack access token. Log into your Slack account and retrieve a token
  from the Slack Web API page.

channel
: Slack message destination. Can be either a #channel or a @username.

username
: CATI bot username.

# EXAMPLES

    ---
    nsuser:
      soundName: Ping
      soundNameFail: Basso
    say:
      voice: Roshi
    espeak:
      voiceName: english-us
    speechsynthesizer:
      voice: Microsoft Roshi Desktop
    bearychat:
      incomingHookURI: 1234567890abcdefg
    hipchat:
      accessToken: 1234567890abcdefg
      room: 1234567890abcdefg
    pushbullet:
      accessToken: 1234567890abcdefg
      deviceIden: 1234567890abcdefg
    pushover:
      userKey: 1234567890abcdefg
      apiToken: 1234567890abcdefg
    pushsafer:
      key: 1234567890abcdefg
    simplepush:
      key: 1234567890abcdefg
      event: 1234567890abcdefg
    slack:
      token: 1234567890abcdefg
      channel: '@roshi619'
      username: cati

# SEE ALSO

cati(1)
