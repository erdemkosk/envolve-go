<p align="center">
  <h1 align="center">Bubbletea-Starter</h1>
  <p align="center">
    <a href="https://github.com/erdemkosk/envolve-go/releases"><img src="https://img.shields.io/github/v/release/knipferrc/bubbletea-starter" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/erdemkosk/envolve-go?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/erdemkosk/envolve-go/actions"><img src="https://img.shields.io/github/workflow/status/knipferrc/bubbletea-starter/Release" alt="Build Status"></a>
  </p>
</p>

## About The Project

A starting point for bubbletea apps

### Built With

- [Go](https://golang.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [Viper](https://github.com/spf13/viper)
- [Cobra](https://github.com/spf13/cobra)

## Installation

### Curl

```sh
curl -sfL https://raw.githubusercontent.com/knipferrc/bubbletea-starter/main/install.sh | sh
```

### Go

```
go install github.com/erdemkosk/envolve-go@latest
```

## Features

- Add your awesome feature list here

## Configuration

A config file will be generated (`bubbletea-starter.yml`) in the config directory of the OS in which the app is ran from. If `XDG_CONFIG_HOME` is set, that will be used instead.

```yml
settings:
  enable_logging: false
  enable_mousewheel: true
```
