# Go Proverbs CLI

Go sample project for a neat littly command line application in Go.

Initially set up for the article for entickler.de [Kommandozeilen-Apps mit Go](https://entwickler.de/go/go-kommandozeilen-apps) (in German).

Features:
- Flags with [POSIX standard format](https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html)
- Binary release with [Goreleaser](https://goreleaser.com/) and [Github Actions](https://github.com/features/actions)
- Usage of [Cobra](https://github.com/spf13/cobra) library (see branch `cobra`)

## Build it

Compile the CLI with:

    go build .


## Run it

Run it and print a Go proverb to the command line with.

    ./progo
