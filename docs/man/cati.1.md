% CATI(1) cati 3.1.0 | CATI Manual
% Jaime Piña
% 2018/03/25

#  NAME

cati - monitor a process and trigger a catification

# SYNOPSIS

cati [flags] [utility [args...]]

# DESCRIPTION

Never sit and wait for some long-running process to finish. CATI can alert you
when it's done. You can receive messages on your computer or phone.

# OPTIONS

-t \<string\>, \--title \<string\>
: Set catification title. Default is utility name.

-m \<string\>, \--message \<string\>
: Set catification message. Default is "Done!".

-b, \--banner
: Trigger a banner catification. This is enabled by default. To disable this
  service, set this flag to false. This will be either `nsuser`, `freedesktop`,
  or `catifyicon` catification, depending on the OS.

-s, \--speech
: Trigger a speech catification. This will be either `say`, `espeak`, or
  `speechsynthesizer` catification, depending on the OS.

-c, \--bearychat
: Trigger a BearyChat catification. This requires `bearychat.incomingHookURI` to
  be set.

-i, \--hipchat
: Trigger a HipChat catification. This requires `hipchat.accessToken` and
  `hipchat.room` to be set.

-p, \--pushbullet
: Trigger a Pushbullet catification. This requires `pushbullet.accessToken` to
  be set.

-o, \--pushover
: Trigger a Pushover catification. This requires `pushover.apiToken` and
  `pushover.userKey` to be set.

-u, \--pushsafer
: Trigger a Pushsafer catification. This requires `pushsafer.key` to be set.

-l, \--simplepush
: Trigger a Simplepush catification. This requires `simplepush.key` to be set.

-k, \--slack
: Trigger a Slack catification. This requires `slack.token` and `slack.channel`
  to be set.

-w <pid>, \--pwatch <pid>
: Monitor a process by PID and trigger a catification when the pid disappears.

-f, \--file
: Path to `cati.yaml` configuration file.

\--verbose
: Enable verbose mode.

-v, \--version
: Print `cati` version and exit.

-h, \--help
: Print `cati` help and exit.

# ENVIRONMENT

* `CATI_DEFAULT`
* `CATI_NSUSER_SOUNDNAME`
* `CATI_NSUSER_SOUNDNAMEFAIL`
* `CATI_SAY_VOICE`
* `CATI_ESPEAK_VOICENAME`
* `CATI_SPEECHSYNTHESIZER_VOICE`
* `CATI_BEARYCHAT_INCOMINGHOOKURI`
* `CATI_HIPCHAT_ACCESSTOKEN`
* `CATI_HIPCHAT_ROOM`
* `CATI_PUSHBULLET_ACCESSTOKEN`
* `CATI_PUSHBULLET_DEVICEIDEN`
* `CATI_PUSHOVER_APITOKEN`
* `CATI_PUSHOVER_USERKEY`
* `CATI_PUSHSAFER_KEY`
* `CATI_SIMPLEPUSH_KEY`
* `CATI_SIMPLEPUSH_EVENT`
* `CATI_SLACK_TOKEN`
* `CATI_SLACK_CHANNEL`
* `CATI_SLACK_USERNAME`


# FILES

If not explicitly set with \--file, then cati will check the following paths,
in the following order.

* ./.cati.yaml
* $XDG_CONFIG_HOME/cati/cati.yaml

# EXAMPLES

Display a catification when `tar` finishes compressing files.

    cati tar -cjf music.tar.bz2 Music/

Add cati after a command, in case you forgot at the beginning.

    clang foo.c -Wall -lm -L/usr/X11R6/lib -lX11 -o bizz; cati

If you already started a command, but forgot to use `cati`, then you can do
this to get catified when that process' PID disappears.

    cati --pwatch $(pgrep docker-machine)

# REPORTING BUGS

Report bugs on GitHub at https://github.com/roshi619/cati/issues.

# SEE ALSO

cati.yaml(5)
