# cati
To monitor a process and trigger notification.


Idea is to alert for any long-running process rather than waiting to finish.

## Services

CATI can send catifications on a number of services like Banner, Hipchat, Slack, Pushover et al


```

## Installation

The `master` branch always contains the latest tagged release.

```shell
# Install the latest version on any platform.
go get -u github.com/roshi619/cati/cmd/cati

```

If you don't want to build from source or install anything extra, just download the latest binary.

```shell
# macOS
curl -L $(curl -s https://api.github.com/repos/roshi619/cati/releases/latest | awk '/browser_download_url/ { print $2 }' | grep 'darwin-amd64' | sed 's/"//g') | tar -xz

# Linux
curl -L $(curl -s https://api.github.com/repos/roshi619/cati/releases/latest | awk '/browser_download_url/ { print $2 }' | grep 'linux-amd64' | sed 's/"//g') | tar -xz
```

Or download with your browser from the [latest release] page.

## Examples

Just put `cati` at the beginning or end of your regular commands. For more
details, checkout the [docs].

Display a catification when `tar` finishes compressing files.

```
cati tar -cjf music.tar.bz2 Music/
```

Add `cati` after a command, in case you forgot at the beginning.

```
clang foo.c -Wall -lm -L/usr/X11R6/lib -lX11 -o bizz; cati
```

If you already started a command, but forgot to use `cati`, then you can do
this to get catified when that process' PID disappears.

```
cati --pwatch $(pgrep docker-machine)
```

You can also press `ctrl+z` after you started a process. This will temporarily
suspend the process, but you can resume it with `cati`.

```
$ dd if=/dev/zero of=foo bs=1M count=2000
^Z
zsh: suspended  dd if=/dev/zero of=foo bs=1M count=2000
$ fg; cati
[1]  + continued  dd if=/dev/zero of=foo bs=1M count=2000
2000+0 records in
2000+0 records out
2097152000 bytes (2.1 GB, 2.0 GiB) copied, 12 s, 175 MB/s
```


[CircleCI]: https://circleci.com/gh/roshi619/cati/tree/master.svg?style=svg
[AppVeyor]: https://ci.appveyor.com/api/projects/status/qc2fgc164786jws6/branch/master?svg=true
[Codecov]: https://codecov.io/gh/roshi619/cati/branch/master/graph/badge.svg
[macOS Banner CATIfication]: https://raw.githubusercontent.com/roshi619/cati/master/docs/screenshots/macos_banner.png
[screenshots]: https://github.com/roshi619/cati/tree/master/docs/screenshots
[latest release]: https://github.com/roshi619/cati/releases/latest
[docs]: https://github.com/roshi619/cati/blob/master/docs/cati.md
