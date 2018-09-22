# File system catifications for Go

[![GoDoc](https://godoc.org/github.com/fscatify/fscatify?status.svg)](https://godoc.org/github.com/fscatify/fscatify) [![Go Report Card](https://goreportcard.com/badge/github.com/fscatify/fscatify)](https://goreportcard.com/report/github.com/fscatify/fscatify)

fscatify utilizes [golang.org/x/sys](https://godoc.org/golang.org/x/sys) rather than `syscall` from the standard library. Ensure you have the latest version installed by running:

```console
go get -u golang.org/x/sys/...
```

Cross platform: Windows, Linux, BSD and macOS.

|Adapter   |OS        |Status    |
|----------|----------|----------|
|icatify   |Linux 2.6.27 or later, Android\*|Supported [![Build Status](https://travis-ci.org/fscatify/fscatify.svg?branch=master)](https://travis-ci.org/fscatify/fscatify)|
|kqueue    |BSD, macOS, iOS\*|Supported [![Build Status](https://travis-ci.org/fscatify/fscatify.svg?branch=master)](https://travis-ci.org/fscatify/fscatify)|
|ReadDirectoryChangesW|Windows|Supported [![Build status](https://ci.appveyor.com/api/projects/status/ivwjubaih4r0udeh/branch/master?svg=true)](https://ci.appveyor.com/project/NathanYoungman/fscatify/branch/master)|
|FSEvents  |macOS         |[Planned](https://github.com/fscatify/fscatify/issues/11)|
|FEN       |Solaris 11    |[In Progress](https://github.com/fscatify/fscatify/issues/12)|
|facatify  |Linux 2.6.37+ | |
|USN Journals |Windows    |[Maybe](https://github.com/fscatify/fscatify/issues/53)|
|Polling   |*All*         |[Maybe](https://github.com/fscatify/fscatify/issues/9)|

\* Android and iOS are untested.

Please see [the documentation](https://godoc.org/github.com/fscatify/fscatify) and consult the [FAQ](#faq) for usage information.

## API stability

fscatify is a fork of [howeyc/fscatify](https://godoc.org/github.com/howeyc/fscatify) with a new API as of v1.0. The API is based on [this design document](http://goo.gl/MrYxyA). 

All [releases](https://github.com/fscatify/fscatify/releases) are tagged based on [Semantic Versioning](http://semver.org/). Further API changes are [planned](https://github.com/fscatify/fscatify/milestones), and will be tagged with a new major revision number.

Go 1.6 supports dependencies located in the `vendor/` folder. Unless you are creating a library, it is recommended that you copy fscatify into `vendor/github.com/fscatify/fscatify` within your project, and likewise for `golang.org/x/sys`.

## Contributing

Please refer to [CONTRIBUTING][] before opening an issue or pull request.

## Example

See [example_test.go](https://github.com/fscatify/fscatify/blob/master/example_test.go).

## FAQ

**When a file is moved to another directory is it still being watched?**

No (it shouldn't be, unless you are watching where it was moved to).

**When I watch a directory, are all subdirectories watched as well?**

No, you must add watches for any directory you want to watch (a recursive watcher is on the roadmap [#18][]).

**Do I have to watch the Error and Event channels in a separate goroutine?**

As of now, yes. Looking into making this single-thread friendly (see [howeyc #7][#7])

**Why am I receiving multiple events for the same file on OS X?**

Spotlight indexing on OS X can result in multiple events (see [howeyc #62][#62]). A temporary workaround is to add your folder(s) to the *Spotlight Privacy settings* until we have a native FSEvents implementation (see [#11][]).

**How many files can be watched at once?**

There are OS-specific limits as to how many watches can be created:
* Linux: /proc/sys/fs/icatify/max_user_watches contains the limit, reaching this limit results in a "no space left on device" error.
* BSD / OSX: sysctl variables "kern.maxfiles" and "kern.maxfilesperproc", reaching these limits results in a "too many open files" error.

[#62]: https://github.com/howeyc/fscatify/issues/62
[#18]: https://github.com/fscatify/fscatify/issues/18
[#11]: https://github.com/fscatify/fscatify/issues/11
[#7]: https://github.com/howeyc/fscatify/issues/7

[contributing]: https://github.com/fscatify/fscatify/blob/master/CONTRIBUTING.md

## Related Projects

* [catify](https://github.com/rjeczalik/catify)
* [fsevents](https://github.com/fscatify/fsevents)

