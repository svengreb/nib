<p align="center"><img src="https://raw.githubusercontent.com/svengreb/nib/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center"><a href="https://github.com/svengreb/nib/releases/latest"><img src="https://img.shields.io/github/release/svengreb/nib.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center">Changelog of the “stylish“, log-level based line printer for human-facing <a href="https://go.dev" target="_blank">Go</a> CLI applications.</p>

<!--lint disable no-duplicate-headings no-duplicate-headings-in-section-->

# 0.2.0

![Release Date: 2020-11-11](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-11-11&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/nib/projects/5) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/nib/milestone/2)

⇅ [Show all commits][gh-compare-tag-v0.1.0_v0.2.0]

## Features

<details>
<summary><strong>Extend nib API to allow reuse of underlying writer</strong> — #21 ⇄ #22 (⊶ 2dac3395)</summary>

↠ To allow to reuse the underlying [io.Writer][go-doc-pkg-type-io#writer] the new `Writer() io.Writer` method has been added to the [`nib.Nib` interface][go-pkg-if-svengreb/nib#nib].
Both implementations the [`pencil`][go-pkg-svengreb/nib/pencil] and [`inkpen`][go-pkg-svengreb/nib/inkpen] package have been adapted to this change.

</details>

<details>
<summary><strong>Update to <code>tmpl-go</code> template repository version 0.5.0</strong> — #25 ⇄ #26 (⊶ 15285fb1)</summary>

↠ Updated to [`tmpl-go` version 0.5.0][gh-tmpl-go-rel-v0.5.0] (including [version 0.4.0][gh-tmpl-go-rel-v0.4.0]) that…

1. …[introduces the initial project documentation][svengreb/tmpl-go#32].
2. …[updates golangci-lint to the currently latest version 1.32.0][svengreb/tmpl-go#21] which introduces new linters like [errorlint][]., [tparallel][]. and [wrapcheck][].
3. …updates to `tmpl` version 0.7.0 ([#25][svengreb/tmpl-go#25] [#34][svengreb/tmpl-go#34]).
   This includes…
   - …a new configuration file for [automated dependency updates and security alerts][svengreb/tmpl#52] with [Dependabot][]. Next to update configurations for the [CI/CD GitHub action workflow][tmpl#cicd] and [Yarn/NPM dependencies][tmpl#node], the file has been extended to support [Go modules][go-doc-mod].
   - …updates to the latest Node.js package dependency & GitHub Action versions.
   - …a [change of the NPM package name to use a namespace][svengreb/tmpl#48] which helps to prevent collisions with already existing NPM packages like [tmpl][npm-tmpl].

</details>

## Improvement

<details>
<summary><strong>Removed <code>pencil.NewFrom(*pencil.Options)</code> function</strong> — #23 ⇄ #24 (⊶ 175478ec)</summary>

↠ The [`pencil.NewFrom(*pencil.Options)`][gh-blob-pencil-2dac3395#l26-27] function was not necessary because…

1. …there is currently no way to get the actual options from a `pencil` instance.
2. …new `pencil` instances with different options can be simply composed by using the variadic parameter of the `pencil.New(pencil.Option...)` function.

Therefore the `pencil.NewFrom(*pencil.Options)` has been removed to simply the package surface.

</details>

<details>
<summary><strong>Updated to latest Go module dependency versions</strong> — #28</summary>

↠ Bumped outdated Go module dependencies to their latest versions:

- #28 (⊶ 76653193) [`github.com/fatih/color`][gh-fatih/color]] from [1.9.0 to 1.10.0][gh-fatih/color-v1.9.0_v1.10.0] — upgrades the `github.com/mattn/go-colorable` and `github.com/mattn/go-isatty` dependencies which include various bug fixes and improvements.

</details>

# 0.1.0

![Release Date: 2020-09-28](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-09-28&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/nib/projects/4) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/nib/milestone/1)

⇅ [Show all commits][repo-compare-tag-init_v0.1.0]

This is the initial release version of _nib_.
The basic project setup, structure and development workflow has been bootstrapped by [the _tmpl-go_ template repository][gh-tmpl-go].

## Features

<details>
<summary><strong>Bootstrap based on "tmpl-go" template repository</strong> — #1 (⊶ 966454c6)</summary>

<p align="center"><img src="https://github.com/svengreb/tmpl-go/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

↠ Bootstrapped the basic project setup, structure and development workflow [from version 0.1.0][gh-tmpl-go-rl-v0.1.0] of the ["tmpl-go" template repository][gh-tmpl-go].
Additionally specific assets like the repository hero image have been replaced and documentations like the _README_ and GitHub issue/PR templates are adjusted.

</details>

<details>
<summary><strong>nib API <code>v0</code></strong> — #2 ⇄ #5 (⊶ 40039091)</summary>

↠ Implemented the API `v0` with the following interfaces and types:

- Ⓘ `Nib` — a log-level based line printer for human-facing messages.
- Ⓣ `Verbosity` — defines the verbosity level of the line printer.

</details>

<details>
<summary><strong>nib.Nib implementation <code>inkpen</code> with color and icon prefix support</strong> — #6 ⇄ #7 (⊶ f5cb3546)</summary>

↠ Implemented a new package `inkpen` that implements the `nib.Nib` API `v0` (designed in #5) with color support, custom prefixes and verbosity level icons for human-facing CLI messages.

The default IO output stream is configurable by accepting a [io.Writer][go-doc-pkg-type-io#writer]. The default writer is [color.Output][gh-blob-fatih/color#output] of the [github.com/fatih/color][gh-fatih/color] which provides automatic TTY and terminal color support detection.

The following [verbosity levels][repo-blob-verbosity-40039091] have been implemented with default icons:

- ⦿ `debug`
- ✕ `error`
- ⭍ `fatal`
- ➜ `info`
- ✓ `success`
- ! `warning`

</details>

<details>
<summary><strong>Repository and package documentations</strong> — #8 ⇄ #9 (⊶ 191456ea)</summary>

↠ Wrote the repository and package documentation which includes information and guides about provided **features**, the currently latest **API** `v0`, general **usage instructions** as well as detailed information about **configurations** like the [**verbosity**][go-pkg-type-svengreb/nib#verbosity], customizable [**icons**][repo-blob-inkpen-default_icons-f5cb3546] and [**prefixes**][go-pkg-met-svengreb/nib/inkpen#setprefixes].

</details>

<details>
<summary><strong>Splitted icon and prefix features into new <code>pencil</code> package</strong> — #18 ⇄ #19 (⊶ 4f97c4c6)</summary>

↠ Before the `inkpen` package provided a writer that was able to print verbosity level icons and prefixes with colors. To allow to use icons and prefixes without colors both features have been moved into a [new `pencil` package][go-pkg-svengreb/nib/pencil] with a [new `pencil.Pencil` type][go-pkg-type-svengreb/nib/pencil#pencil]. The [`inkpen.Inkpen` type][go-pkg-type-svengreb/nib/inkpen#inkpen] has been changed to compose this new type and additionally allows to color the output.

In order to make it easier to configure both types, new package specific options has been implemented that are available via variadic parameters to the constructor functions of these types. This reduces the required code to configure a writer and also provides a safer usage with _Goroutines_.

</details>

## Improvements

<details>
<summary><strong>Optimized OS and Go version matrix strategy for CI workflow</strong> — #12 ⇄ #13 (⊶ 2fca9af2)</summary>

↠ Before the CI workflow used a matrix strategy to run the `lint-node`, `lint-go` and `test` jobs, but this also included steps that were not necessary for this repository. This has been improved to make the workflow run faster by avoiding unnecessary steps:

- The `lint-node` job has been changed to only run on the [currently latest stable Node version `14.x`][nodejs/node-cl-v14] only on _Linux_ because this repository is not focused on JavaScript but only runs Node based tools to lint other files within this repository.
- The `lint-go` job has been changed to only run on the [currently latest stable Go version `1.15.x`][go-doc-rln-1.15] only on _Linux_ because `golangci-lint` doesn't care about the _Go_ version and OS it runs on but only statically checks the source code.
- The `test` job has been changed to only run on the [currently latest stable Go version `1.15.x`][go-doc-rln-1.15] and only _Linux_ and _Windows_ while _macOS_ is not necessary for this repository because there is no _macOS_ specific code.

These changes also help to keep the required GitHub Action run minutes for the account of this repository as small as possible without wasting resources for unnecessary tasks.

</details>

<details>
<summary><strong>Properly merge configured inkpen icon map</strong> — #16 ⇄ #17 (⊶ 17c8d6d3)</summary>

↠ Before the `inkpen.SetIcons(map[nib.Verbosity]Icon)` method has overridden a `nib.Verbosity` entry instead of merging the `ColorFunc` and `Value` fields causing unexpected behavior like uncolored output or missing icons. This has been changed to check for the _zero value_ of the fields and merge them with the already configured struct.

</details>

## Tasks

<details>
<summary><strong>Adapted to "tmpl-go" repository template version 0.1.1</strong> — #3 ⇄ #4 (⊶ 0998f7e4)</summary>

↠ Adapted to ["tmpl-go" version 0.1.1][gh-tmpl-go-rl-v0.1.1] which includes an [important bug fix][svengreb/tmpl-go#7] regarding the golangci-lint YAML configuration for the [go-header][] linter that caused golangci-lint to fail.

</details>

<details>
<summary><strong>Adapted to "tmpl-go" template repository version 0.2.0</strong> — #10 ⇄ #11 (⊶ d9268f0c)</summary>

↠ Adapted to ["tmpl-go" version 0.2.0][gh-tmpl-go-rl-v0.2.0] which includes a [basic setup for testing in the CI workflow][svengreb/tmpl-go#11] using _Go_'s official `go test` command with enabled coverage and race detector.

</details>

<details>
<summary><strong>Adapted to "tmpl-go" template repository version 0.3.0</strong> — #14 ⇄ #15 (⊶ 7ddcd614)</summary>

↠ Adapted to ["tmpl-go" version 0.3.0][svengreb/tmpl-go-rel-v0.3.0] which includes a [optimized run configuration for the CI workflow][svengreb/tmpl-go#18] that helps to improve the performance through more fine grained configurations:

- Only runs on pushes to the `main` branch.
- Only runs on pushes for `v*` tags.
- Always runs for pushes to PRs.

</details>

<!--
+------------------+
+ Formatting Notes +
+------------------+

The `<summary />` tag must be separated with a blank line from the actual item content paragraph,
otherwise Markdown elements are not parsed and rendered!

+------------------+
+ Symbol Reference +
+------------------+
↠ (U+21A0): Start of a log section description
— (U+2014): Separator between a log section title and the metadata
⇄ (U+21C4): Separator between a issue ID and pull request ID in a log metadata
⊶ (U+22B6): Icon prefix for the short commit SHA checksum in a log metadata
⇅ (U+21C5): Icon prefix for the link of the Git commit history comparison on GitHub
-->

<!--lint disable final-definition-->

<!-- Shared -->

[gh-fatih/color]: https://github.com/fatih/color

<!-- Base Links -->

[go-pkg-svengreb/nib/pencil]: https://pkg.go.dev/github.com/svengreb/nib/pencil
[go-doc-pkg-type-io#writer]: https://golang.org/pkg/io/#Writer

<!-- v0.1.0 -->

[gh-blob-fatih/color#output]: https://github.com/fatih/color/blob/fccafd9e876be44d0d7b380a3b03aeb661c1e231/color.go#L25
[gh-tmpl-go-rl-v0.1.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.1.0
[gh-tmpl-go-rl-v0.1.1]: https://github.com/svengreb/tmpl-go/releases/tag/v0.1.1
[gh-tmpl-go-rl-v0.2.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.2.0
[gh-tmpl-go]: https://github.com/svengreb/tmpl-go
[go-doc-rln-1.15]: https://golang.org/doc/go1.15
[go-header]: https://github.com/denis-tingajkin/go-header
[go-pkg-met-svengreb/nib/inkpen#setprefixes]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen.SetPrefixes
[go-pkg-type-svengreb/nib/inkpen#inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen#Inkpen
[go-pkg-type-svengreb/nib/pencil#pencil]: https://pkg.go.dev/github.com/svengreb/nib/pencil#Pencil
[go-pkg-type-svengreb/nib#verbosity]: https://pkg.go.dev/github.com/svengreb/nib#Verbosity
[nodejs/node-cl-v14]: https://github.com/nodejs/node/blob/master/doc/changelogs/CHANGELOG_V14.md
[repo-blob-inkpen-default_icons-f5cb3546]: https://github.com/svengreb/nib/blob/f5cb3546efb0efa6fa25ded6dec96fcbf83e9857/inkpen/inkpen.go#L29-L37
[repo-blob-verbosity-40039091]: https://github.com/svengreb/nib/blob/400390915ca497add7c43aa1ff232a08c8501081/verbosity.go
[repo-compare-tag-init_v0.1.0]: https://github.com/svengreb/nib/compare/4ca4af20...v0.1.0
[svengreb/tmpl-go-rel-v0.3.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.3.0
[svengreb/tmpl-go#11]: https://github.com/svengreb/tmpl-go/pull/11
[svengreb/tmpl-go#18]: https://github.com/svengreb/tmpl-go/issues/18
[svengreb/tmpl-go#7]: https://github.com/svengreb/tmpl-go/pull/7

<!-- v0.2.0 -->

[dependabot]: https://dependabot.com
[errorlint]: https://github.com/polyfloyd/go-errorlint
[gh-blob-pencil-2dac3395#l26-27]: https://github.com/svengreb/nib/blob/2dac3395/pencil/pencil.go#L26-L27
[gh-compare-tag-v0.1.0_v0.2.0]: https://github.com/svengreb/nib/compare/v0.1.0...v0.2.0
[gh-fatih/color-v1.9.0_v1.10.0]: https://github.com/fatih/color/compare/v1.9.0...v1.10.0
[gh-tmpl-go-rel-v0.4.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.4.0
[gh-tmpl-go-rel-v0.5.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.5.0
[go-doc-mod]: https://golang.org/ref/mod
[go-pkg-if-svengreb/nib#nib]: https://pkg.go.dev/github.com/svengreb/nib#Nib
[go-pkg-svengreb/nib/inkpen]: https://pkg.go.dev/github.com/svengreb/nib/inkpen
[npm-tmpl]: https://www.npmjs.com/package/tmpl
[svengreb/tmpl-go#21]: https://github.com/svengreb/tmpl-go/pull/21
[svengreb/tmpl-go#25]: https://github.com/svengreb/tmpl-go/issues/25
[svengreb/tmpl-go#32]: https://github.com/svengreb/tmpl-go/pull/32
[svengreb/tmpl-go#34]: https://github.com/svengreb/tmpl-go/issues/34
[svengreb/tmpl#48]: https://github.com/svengreb/tmpl/issues/48
[svengreb/tmpl#52]: https://github.com/svengreb/tmpl/issues/52
[tmpl#cicd]: https://github.com/svengreb/tmpl#cicd-action-workflow
[tmpl#node]: https://github.com/svengreb/tmpl#nodejs-yarn-and-npm
[tparallel]: https://github.com/moricho/tparallel
[wrapcheck]: https://github.com/tomarrell/wrapcheck
