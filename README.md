<p align="center"><img src="https://github.com/svengreb/nib/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

<p align="center"><a href="https://github.com/svengreb/nib/releases/latest" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/svengreb/nib/blob/main/CHANGELOG.md" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=Changelog&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/nib" target="_blank"><img src="https://img.shields.io/github/v/release/svengreb/nib.svg?style=flat-square&label=GoDoc&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center"><a href="https://github.com/svengreb/nib/actions?query=workflow%3Aci" target="_blank"><img src="https://img.shields.io/github/workflow/status/svengreb/nib/ci.svg?style=flat-square&label=CI&logo=github&logoColor=eceff4&colorA=4c566a"/></a> <a href="https://codecov.io/gh/svengreb/nib" target="_blank"><img src="https://img.shields.io/codecov/c/github/svengreb/nib/main.svg?style=flat-square&label=Coverage&logo=codecov&logoColor=eceff4&colorA=4c566a"/></a></p>

<p align="center"><a href="https://golang.org/doc/effective_go.html#formatting" target="_blank"><img src="https://img.shields.io/static/v1?style=flat-square&label=Go%20Style%20Guide&message=gofmt&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/arcticicestudio/styleguide-markdown/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-markdown.svg?style=flat-square&label=Markdown%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzOSIgaGVpZ2h0PSIzOSIgdmlld0JveD0iMCAwIDM5IDM5Ij48cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNEOERFRTkiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBkPSJNMS41IDEuNWgzNnYzNmgtMzZ6Ii8%2BPHBhdGggZmlsbD0iI0Q4REVFOSIgZD0iTTIwLjY4MyAyNS42NTVsNS44NzItMTMuNDhoLjU2Nmw1Ljg3MyAxMy40OGgtMS45OTZsLTQuMTU5LTEwLjA1Ni00LjE2MSAxMC4wNTZoLTEuOTk1em0tMi42OTYgMGwtMTMuNDgtNS44NzJ2LS41NjZsMTMuNDgtNS44NzJ2MS45OTVMNy45MzEgMTkuNWwxMC4wNTYgNC4xNnoiLz48L3N2Zz4%3D"/></a> <a href="https://github.com/arcticicestudio/styleguide-git/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-git.svg?style=flat-square&label=Git%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=git"/></a></p>

<p align="center">A “stylish“, log-level based line printer for human-facing <a href="https://go.dev" target="_blank">Go</a> CLI applications.</p>

## Features

The [`pencil` package][go-pkg-svengreb/nib/pencil] provides [`pencil.Pencil`][go-pkg-type-svengreb/nib/pencil#pencil] that implements [`nib.Nib`][go-pkg-type-svengreb/nib#nib] with

- …custom prefixes
- …configurable verbosity level icons

The [`inkpen` package][go-pkg-svengreb/nib/inkpen] composes `pencil.Pencil` and additionally provides…

- …colored output
- …automatic TTY and cross-platform terminal color support detection

Please note that this package has mainly been created for my personal use in mind to avoid copying source code between my CLI based projects that require a line printer for human-facing messages. The default configurations might not fit your needs, but the `pencil.Pencil` and `inkpen.Inkpen` implementations of the `nib.Nib` interface have been designed so that they can be flexibly adapted to different use cases and environments.

## API

The [`nib` package][go-pkg-svengreb/nib] provides the currently latest API `v0`.
The [`nib.Nib` interface][go-pkg-type-svengreb/nib#nib] consists of six functions that allow to print a formatted message with different [verbosity levels][go-pkg-type-svengreb/nib#verbosity]:

- `Compile(v Verbosity, format string, args ...interface{}) string` — compiles a message for the verbosity level using the given format and arguments..
- `Debugf(format string, args ...interface{})` — writes a message with `debug` verbosity level for the given format and arguments.
- `Errorf(format string, args ...interface{})` — writes a message with `error` verbosity level for the given format and arguments.
- `Fatalf(format string, args ...interface{})` — writes a message with `fatal` verbosity level for the given format and arguments.
- `Infof(format string, args ...interface{})` — writes a message with `info` verbosity level for the given format and arguments.
- `Successf(format string, args ...interface{})` — writes a message with `success` verbosity level for the given format and arguments.
- `Warnf(format string, args ...interface{})` — writes a message with `warn` verbosity level for the given format and arguments.

The [`pencil` package][go-pkg-svengreb/nib/pencil] implements this interface including [features like custom prefixes and verbosity level icons](#features).
The [`inkpen` package][go-pkg-svengreb/nib/inkpen] composes [`pencil.Pencil`][go-pkg-type-svengreb/nib/pencil#pencil] and additionally [comes with additional features like colored output including automatic TTY and cross-platform terminal color support detection](#features).

For more details about the API, available packages and types please see the [GoDoc reference documentation][go-pkg-svengreb/nib].

## Usage

In addition to the possibility of implementing the [`nib.Nib` interface][go-pkg-type-svengreb/nib#nib] yourself the [`pencil.Pencil` type][go-pkg-type-svengreb/nib/pencil#pencil] can be used:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil
	pen := pencil.New()
	// ... or inkpen with default configurations where nib.InfoVerbosity is the default verbosity level.
	ink := inkpen.New()

	// Print a message with "info" level.
	fruit := "coconut"
	pen.Infof("My favorite fruits are %ss", fruit)
	ink.Infof("My favorite fruits are %ss", fruit)
}
```

<!--lint enable no-tabs-->

## Configuration

Both types [`pencil.Pencil` type][go-pkg-type-svengreb/nib/pencil#pencil] [`inkpen.Inkpen` type][go-pkg-type-svengreb/nib/inkpen#inkpen] were designed so that they can be flexibly adapted to different use cases and environments.

To customize a `pencil.Pencil` the [`New` function][go-pkg-func-svengreb/nib/pencil#new] accepts one or more [`pencil.Option`][go-pkg-type-svengreb/nib/pencil#option].
A `inkpen.Inkpen` can also be customized with one or more `pencil.Option` by passing them to the [`inkpen.WithPencilOptions(pencilOpts ...pencil.Option) Option` function][go-pkg-met-svengreb/nib/inkpen#eithpenciloptions] via the [`New` function][go-pkg-func-svengreb/nib/inkpen#new].

### Verbosity

The default verbosity level of `pencil.Pencil` is [`nib.InfoVerbosity`][go-pkg-var-svengreb/nib#infoverbosity] that prints messages with “info“ scope.
You can adjust it to any other level through the [`pencil.WithVerbosity(verbosity nib.Verbosity) Option`][go-pkg-met-svengreb/nib/pencil#withverbosity] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New(pencil.WithVerbosity(nib.DebugVerbosity))
	// ...or inkpen with default configurations and set the verbosity level from "info" to "debug" scope.
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithVerbosity(nib.DebugVerbosity)))

	pen.Debugf("Raspberries are my second favorite fruit")
	ink.Debugf("Raspberries are my second favorite fruit")
}
```

<!--lint enable no-tabs-->

Note that [`nib.Verbosity`][go-pkg-type-svengreb/nib#verbosity] implements [`encoding.TextUnmarshaler`][go-doc-encoding#textnnmarshaler] to allow to unmarshal a textual representation of the level or get the `nib.Verbosity` based on the level scope name:

<!--lint disable no-tabs-->

```go
package main

import (
	"fmt"
	"os"

	"github.com/svengreb/nib"
)

func main() {
	// Get the textual representation of the verbosity level.
	fmt.Println(nib.InfoVerbosity.String()) // "info"

	// Get the nib.Verbosity from the given verbosity level scope name.
	// The function returns an non-nil error when the given name is not valid.
	v, err := nib.ParseVerbosity("error")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Verbosity: %s\n", v.String()) // "error"
}
```

<!--lint enable no-tabs-->

To check whether the verbosity level is enabled, the [`*pencil.Pencil.Enabled(v nib.Verbosity) bool`][go-pkg-met-svengreb/nib/pencil#enabled] or [`*inkpen.Inkpen.Enabled(v nib.Verbosity) bool`][go-pkg-met-svengreb/nib/inkpen#enabled] methods can be used:

<!--lint disable no-tabs-->

```go
package main

import (
	"fmt"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil
	pen := pencil.New()
	// ... or inkpen...
	ink := inkpen.New()

	// ...and check whether the verbosity level is enabled.
	fmt.Println(pen.Enabled(nib.DebugVerbosity)) // false
	fmt.Println(pen.Enabled(nib.ErrorVerbosity)) // true
	fmt.Println(ink.Enabled(nib.FatalVerbosity)) // true
}
```

<!--lint enable no-tabs-->

### Icons

The default icons are [Unicode characters][wikip-unicode] which are supported by almost all terminals nowadays, but there might be cases where you want to use simple [ASCII characters][wikip-ascii] instead.
You can change any verbosity level icon through the [`pencil.WithIcons(icons map[nib.Verbosity]rune) Option`][go-pkg-met-svengreb/nib/pencil#withicons] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New(pencil.WithIcons(map[nib.Verbosity]rune{nib.InfoVerbosity: '>'}))
	// ...or inkpen, that uses Unicode characters as verbosity level icons by default, and change the icon for the "info"
	// verbosity level scope.
	ink := inkpen.New(inkpen.WithPencilOptions(
		pencil.WithIcons(map[nib.Verbosity]rune{nib.InfoVerbosity: '>'}),
	))

	pen.Infof("Cashew nuts can be combined very well with yogurt")
	ink.Infof("Cashew nuts can be combined very well with yogurt")
}

```

<!--lint enable no-tabs-->

`inkpen.Inkpen` allows to customize the color of icons through the [`inkpen.WithIconColorFuncs(iconColorFuncs map[nib.Verbosity]IconColorFunc) Option`][go-pkg-met-svengreb/nib/inkpen#withiconcolorfuncs] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/fatih/color"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
)

func main() {
	// Create a new inkpen and change the color function for the "info" verbosity level scope to use green instead of
	// blue as foreground color.
	ink := inkpen.New(inkpen.WithIconColorFuncs(map[nib.Verbosity]inkpen.IconColorFunc{
		nib.InfoVerbosity: color.New(color.FgGreen).Sprintf,
	}))

	ink.Infof("Almonds taste very good with vanilla")
}
```

<!--lint enable no-tabs-->

To disable verbosity level icons to be printed at all use the [`pencil.UseIcons(useIcons bool) Option`][go-pkg-met-svengreb/nib/pencil#useicons] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New(pencil.UseIcons(false))
	// ...or inkpen and disable all verbosity level icons.
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.UseIcons(false)))

	pen.Errorf("Cane sugar is not necessary for a delicious yogurt")
	ink.Errorf("Cane sugar is not necessary for a delicious yogurt")
}

```

<!--lint enable no-tabs-->

Note that the provided map of icons gets merged with the default map in order to ensure there are no missing icons.
See the [`pencil` package][go-pkg-svengreb/nib/pencil] to get an overview of the icon characters that are used by default.

To check if verbosity level icons are enabled, the [`*pencil.Pencil.IconsEnabled() bool`][go-pkg-met-svengreb/nib/pencil#iconsenabled] or [`*inkpen.Inkpen.IconsEnabled() bool`][go-pkg-met-svengreb/nib/inkpen#iconsenabled] methods can be used:

<!--lint disable no-tabs-->

```go
package main

import (
	"fmt"

	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New()
	// ...and check whether verbosity level icons are enabled.
	fmt.Println(pen.IconsEnabled()) // true
}
```

### Prefixes

By default only the verbosity level icons are printed as prefix before the given message format and arguments.
You can add any amount of custom prefixes through the [`pencil.WithPrefixes(prefixes ...string) Option`][go-pkg-met-svengreb/nib/pencil#withprefixes] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New(pencil.WithPrefixes("[fruit mixer]"))
	// ... or inkpen and add a custom prefix that gets placed before the actual message.
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithPrefixes("[fruit mixer]")))

	pen.Infof("Strawberries are also a very tasty ingredient for low-fat quark")
	ink.Infof("Strawberries are also a very tasty ingredient for low-fat quark")
}
```

### Writer

By default `pencil.Pencil` uses [`os.Stderr`][go-doc-os#pkg_vars] while `inkpen.Inkpen` uses `color.Output` which in turn is a exported variable that makes use of [github.com/mattn/go-colorable][mattn/go-colorable], a package for colored TTY output on multiple platforms.
You can use any [`io.Writer`][go-doc-type-io#writer] through the [`pencil.WithWriter(writer io.Writer) Option`][go-pkg-met-svengreb/nib/pencil#withicons] function:

<!--lint disable no-tabs-->

```go
package main

import (
	"os"

	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"
)

func main() {
	// Create a new pencil...
	pen := pencil.New(pencil.WithWriter(os.Stderr))
	// ... or inkpen and change the output writer to use the OS error stream.
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithWriter(os.Stderr)))

	pen.Errorf("Blueberries mixed with raspberries and yoghurt are a delicious dream")
	ink.Errorf("Blueberries mixed with raspberries and yoghurt are a delicious dream")
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
[go-doc-os#pkg_vars]: https://golang.org/pkg/os/#pkg-variables
[go-doc-type-io#writer]: https://golang.org/pkg/io/#Writer
[go-pkg-func-svengreb/nib/inkpen#new]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#New
[go-pkg-func-svengreb/nib/pencil#new]: https://pkg.go.dev/github.com/svengreb/nib/pencil#New
[go-pkg-met-svengreb/nib/inkpen#eithpenciloptions]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.WithPencilOptions
[go-pkg-met-svengreb/nib/inkpen#enabled]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.Enabled
[go-pkg-met-svengreb/nib/inkpen#iconsenabled]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.IconsEnabled
[go-pkg-met-svengreb/nib/inkpen#withiconcolorfuncs]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.WithIconColorFuncs
[go-pkg-met-svengreb/nib/pencil#enabled]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.Enabled
[go-pkg-met-svengreb/nib/pencil#iconsenabled]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.IconsEnabled
[go-pkg-met-svengreb/nib/pencil#useicons]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.UseIcons
[go-pkg-met-svengreb/nib/pencil#withicons]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.WithIcons
[go-pkg-met-svengreb/nib/pencil#withprefixes]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.WithPrefixes
[go-pkg-met-svengreb/nib/pencil#withverbosity]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil.WithVerbosity
[go-pkg-svengreb/nib]: https://pkg.go.dev/github.com/svengreb/nib
[go-pkg-svengreb/nib/inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen
[go-pkg-svengreb/nib/pencil]: https://pkg.go.dev/github.com/svengreb/nib/pencil
[go-pkg-type-svengreb/nib/inkpen#inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen
[go-pkg-type-svengreb/nib/pencil#option]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Option
[go-pkg-type-svengreb/nib/pencil#pencil]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil
[go-pkg-type-svengreb/nib#nib]: https://pkg.go.dev/github.com/svengreb/nib#Nib
[go-pkg-type-svengreb/nib#verbosity]: https://pkg.go.dev/github.com/svengreb/nib#Verbosity
[go-pkg-var-svengreb/nib#infoverbosity]: https://pkg.go.dev/github.com/svengreb/nib#InfoVerbosity
[mattn/go-colorable]: https://github.com/mattn/go-colorable
[wikip-ascii]: https://en.wikipedia.org/wiki/ASCII
[wikip-unicode]: https://en.wikipedia.org/wiki/Unicode
