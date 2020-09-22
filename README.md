<p align="center"><img src="https://github.com/svengreb/nib/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

<p align="center"><a href="https://github.com/svengreb/nib/releases/latest" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/svengreb/nib/blob/main/CHANGELOG.md" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=Changelog&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/nib" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=GoDoc&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center"><a href="https://github.com/svengreb/nib/actions?query=workflow%3Aci" target="_blank"><img src="https://img.shields.io/github/workflow/status/svengreb/nib/ci.svg?style=flat-square&label=CI&logo=github&logoColor=eceff4&colorA=4c566a"/></a> <a href="https://codecov.io/gh/svengreb/nib" target="_blank"><img src="https://img.shields.io/codecov/c/github/svengreb/nib/main.svg?style=flat-square&label=Coverage&logo=codecov&logoColor=eceff4&colorA=4c566a"/></a></p>

<p align="center"><a href="https://golang.org/doc/effective_go.html#formatting" target="_blank"><img src="https://img.shields.io/static/v1?style=flat-square&label=Go%20Style%20Guide&message=gofmt&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/arcticicestudio/styleguide-markdown/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-markdown.svg?style=flat-square&label=Markdown%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzOSIgaGVpZ2h0PSIzOSIgdmlld0JveD0iMCAwIDM5IDM5Ij48cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNEOERFRTkiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBkPSJNMS41IDEuNWgzNnYzNmgtMzZ6Ii8%2BPHBhdGggZmlsbD0iI0Q4REVFOSIgZD0iTTIwLjY4MyAyNS42NTVsNS44NzItMTMuNDhoLjU2Nmw1Ljg3MyAxMy40OGgtMS45OTZsLTQuMTU5LTEwLjA1Ni00LjE2MSAxMC4wNTZoLTEuOTk1em0tMi42OTYgMGwtMTMuNDgtNS44NzJ2LS41NjZsMTMuNDgtNS44NzJ2MS45OTVMNy45MzEgMTkuNWwxMC4wNTYgNC4xNnoiLz48L3N2Zz4%3D"/></a> <a href="https://github.com/arcticicestudio/styleguide-git/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-git.svg?style=flat-square&label=Git%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=git"/></a></p>

<p align="center">A “stylish“, log-level based line printer for human-facing <a href="https://go.dev" target="_blank">Go</a> CLI applications.</p>

_nib_ is a log-level based line printer for human-facing Go CLI applications.

## Features

The [`inkpen` package][go-pkg-svengreb/nib/inkpen] provides [`inkpen.Inkpen`][go-pkg-type-svengreb/nib/inkpen#inkpen] that implements [`nib.Nib`][go-pkg-type-svengreb/nib#nib] with…

- …colored output
- …automatic TTY detection
- …custom prefixes
- …configurable verbosity level icons

Please note that this package has mainly been created with my personal use in mind to avoid copying source code between my CLI based projects that require a line printer for human-facing messages. The default configurations might not fit your needs, but the provided `inkpen.Inkpen` implementation of the `nib.Nib` interface was designed so that it can be flexibly adapted to different use cases and environments.

## API

The [`nib` package][go-pkg-svengreb/nib] provides the currently latest API `v0`.
The [`nib.Nib` interface][go-pkg-type-svengreb/nib#nib] consists of six functions that allow to log a formatted message with different [verbosity log levels][go-pkg-type-svengreb/nib#verbosity]:

- `Debugf(format string, args ...interface{})` — writes a message with `debug` verbosity level for the given format and arguments.
- `Errorf(format string, args ...interface{})` — writes a message with `error` verbosity level for the given format and arguments.
- `Fatalf(format string, args ...interface{})` — writes a message with `fatal` verbosity level for the given format and arguments.
- `Infof(format string, args ...interface{})` — writes a message with `info` verbosity level for the given format and arguments.
- `Successf(format string, args ...interface{})` — writes a message with `success` verbosity level for the given format and arguments.
- `Warnf(format string, args ...interface{})` — writes a message with `warn` verbosity level for the given format and arguments.

The [`inkpen` package][go-pkg-svengreb/nib/inkpen] is the default implementation of this interface that [comes with additional features like colorized output and log level icons](#features).

For more details about the API, available packages and types please see the [GoDoc reference documentation][go-pkg-svengreb/nib].

## Usage

In addition to the possibility of implementing the [`nib.Nib` interface][go-pkg-type-svengreb/nib#nib] yourself the [`inkpen.Inkpen` type][go-pkg-type-svengreb/nib/inkpen#inkpen] can be used:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with default configurations.
	// By default the nib.InfoVerbosity log level is set.
	ink := inkpen.NewDefault()

	// Print a message with "info" level.
	fruit := "coconut"
	ink.Infof("My favorite fruits are %ss", fruit)
}
```

<!--lint enable no-tabs-->

To initialize a new `inkpen.Inkpen` with another default verbosity log level and writer the [`inkpen.New(v nib.Verbosity, w io.Writer) *Inkpen` function][go-pkg-func-svengreb/nib/inkpen#new] can be used:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with "error" as default log level verbosity and the OS error output stream as writer.
	ink := inkpen.New(nib.ErrorVerbosity, os.Stderr)
	ink.Errorf("Cane sugar is not necessary for a delicious yogurt")
}
```

<!--lint enable no-tabs-->

## Configuration

The [`inkpen.Inkpen` type][go-pkg-type-svengreb/nib/inkpen#inkpen] was designed so that it can be flexibly adapted to different use cases and environments.

### Verbosity

The default log level is [`nib.InfoVerbosity`][go-pkg-var-svengreb/nib#infoverbosity] that prints messages with “info“ scope.
You can adjust it to any other level through the [`inkpen.SetVerbosity(v nib.Verbosity)`][go-pkg-met-svengreb/nib/inkpen#setverbosity] method:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with default configurations and set the log level from "info" to "debug" scope.
	ink := inkpen.NewDefault()
	ink.SetVerbosity(nib.DebugVerbosity)
	ink.Debugf("Raspberries are my second favorite fruit")
}
```

<!--lint enable no-tabs-->

Note that [`nib.Verbosity`][go-pkg-type-svengreb/nib#verbosity] implements [`encoding.TextUnmarshaler`][go-doc-encoding#textnnmarshaler] to allow to unmarshal a textual representation of the log level or get the `nib.Verbosity` based on the log level scope name:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Get the textual representation of the log level verbosity.
	fmt.Println(nib.InfoVerbosity.String()) // "info"

	// Get the nib.Verbosity from the given log level scope name.
	// The function returns an non-nil error when the given name is not valid.
	v, err := nib.ParseVerbosity("error")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Verbosity: %s", v.String()) // "error"
}
```

<!--lint enable no-tabs-->

### Icons

The [default icons][repo-blob-inkpen-default_icons] are [Unicode characters][wikip-unicode] which are supported by almost all terminals nowadays, but there might be cases where you want to use simple [ASCII characters][wikip-ascii] instead.
You can change any log level [`inkpen.Icon`][go-pkg-type-svengreb/nib/inkpen#icon] through the [`inkpen.SetIcons(icons map[nib.Verbosity]inkpen.Icon)`][go-pkg-met-svengreb/nib/inkpen#seticons] method:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with default configurations.
	// By default Unicode characters are used as log level icons.
	ink := inkpen.NewDefault()

	// Change the icon for the "info" log level scope.
	ink.SetIcons(map[nib.Verbosity]inkpen.Icon{
		nib.DebugVerbosity: {
			Value:     '>',
		},
	})
  ink.Infof("Cashew nuts can be combined very well with yogurt")
}
```

<!--lint enable no-tabs-->

You can also customize the color of an icon by changing the default function of the [`inkpen.Icon.ColorFunc` field][go-pkg-type-svengreb/nib/inkpen#icon]:

<!--lint disable no-tabs-->

```go
package main

import (
  "github.com/fatih/color"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
  // Create a new inkpen with default configurations.
	ink := inkpen.NewDefault()

	// Change the color function for the "info" log level scope to use green instead of blue as foreground color.
	ink.SetIcons(map[nib.Verbosity]inkpen.Icon{
    nib.InfoVerbosity: {
      ColorFunc: color.New(color.FgGreen).Sprintf,
			Value: '➜',
		},
	})
	ink.Infof("Almonds taste very good with vanilla")
}
```

<!--lint enable no-tabs-->

To disable log level icons to be printed at all set the `inkpen.Inkpen.WithIcons` field to `false`:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with default configurations.
	ink := inkpen.NewDefault()

	// Disable all log level icons.
	ink.WithIcon = false
}
```

<!--lint enable no-tabs-->

Note that the provided map of icons gets merged with the default map in order to ensure there are no missing icons.
See the [`inkpen` package][go-pkg-svengreb/nib/inkpen] to get an overview of the icon characters that are used by default.

### Prefixes

By default only the log level icons are printed as prefix before the given message format and arguments.
You can add amount of custom prefixes through the [`inkpen.SetPrefixes(p ...string)`][go-pkg-met-svengreb/nib/inkpen#setprefixes] method:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen with default configurations.
	ink := inkpen.NewDefault()

	// Add a custom prefix that gets placed before the actual message.
	ink.SetPrefixes("[fruitmixer]")
	ink.Infof("Strawberries are also a very tasty ingredient for low-fat quark")
}
```

<!--lint enable no-tabs-->

## Contributing

_nib_ is an open source project and contributions are always welcome!

There are many ways to contribute, from [writing- and improving documentation and tutorials][contrib-guide-docs], [reporting bugs][contrib-guide-bugs], [submitting enhancement suggestions][contrib-guide-enhance] that can be added to _nib_ by [submitting pull requests][contrib-guide-pr].

Please take a moment to read the [contributing guide][contrib-guide] to learn about the development process, the [styleguides][contrib-guide-styles] to which this project adheres as well as the [branch organization][contrib-guide-branching] and [versioning][contrib-guide-versioning] model.

The guide also includes information about [minimal, complete, and verifiable examples][contrib-guide-mcve] and other ways to contribute to the project like [improving existing issues][contrib-guide-impr-issues] and [giving feedback on issues and pull requests][contrib-guide-feedback].

<p align="center">Copyright &copy; 2019-present <a href="https://www.svengreb.de" target="_blank">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/nib/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

[contrib-guide-branching]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#branch-organization
[contrib-guide-bugs]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#bug-reports
[contrib-guide-docs]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#documentations
[contrib-guide-enhance]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#enhancement-suggestions
[contrib-guide-feedback]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#give-feedback-on-issues-and-pull-requests
[contrib-guide-impr-issues]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#improve-issues
[contrib-guide-mcve]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#mcve
[contrib-guide-pr]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#pull-requests
[contrib-guide-styles]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#styleguides
[contrib-guide-versioning]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md#versioning
[contrib-guide]: https://github.com/svengreb/nib/blob/main/CONTRIBUTING.md
[go-doc-encoding#textnnmarshaler]: https://golang.org/pkg/encoding/#TextUnmarshaler
[go-pkg-func-svengreb/nib/inkpen#new]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#New
[go-pkg-met-svengreb/nib/inkpen#seticons]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.SetIcons
[go-pkg-met-svengreb/nib/inkpen#setprefixes]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.SetPrefixes
[go-pkg-met-svengreb/nib/inkpen#setverbosity]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.SetVerbosity
[go-pkg-svengreb/nib]: https://pkg.go.dev/github.com/svengreb/nib
[go-pkg-svengreb/nib/inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen
[go-pkg-type-svengreb/nib/inkpen#icon]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Icon
[go-pkg-type-svengreb/nib/inkpen#inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen
[go-pkg-type-svengreb/nib#nib]: https://pkg.go.dev/github.com/svengreb/nib#Nib
[go-pkg-type-svengreb/nib#verbosity]: https://pkg.go.dev/github.com/svengreb/nib#Verbosity
[go-pkg-var-svengreb/nib#infoverbosity]: https://pkg.go.dev/github.com/svengreb/nib#InfoVerbosity
[repo-blob-inkpen-default_icons]: https://github.com/svengreb/nib/blob/f5cb3546efb0efa6fa25ded6dec96fcbf83e9857/inkpen/inkpen.go#L29-L37
[wikip-ascii]: https://en.wikipedia.org/wiki/ASCII
[wikip-unicode]: https://en.wikipedia.org/wiki/Unicode
