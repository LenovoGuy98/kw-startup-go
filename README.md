# Kindworks Startup Application

A simple GUI application for Kindworks laptops that provides easy access to system information, the Kindworks website, and documentation.

## Features

- Displays a welcome message and Kindworks logo.
- Provides a link to the Kindworks website.
- Displays system information like OS, architecture, username, hostname, memory, and disk space.
- Provides links to open local documentation files.

## Requirements

- Go
- goGtk library and its dependencies.

```bash
sudo apt-get install libgtk-3-dev libglib2.0-dev
```


## Installation

To install the dependencies, run the following command in the `kw-startup-go` directory:

```bash
go get github.cogo get -u github.com/gotk3/gotk3/gtk@masterm/gotk3/gotk3/gtk

go mod tidy
```

## Usage

To run the application, execute the following command in the `kw-startup-go` directory:

```bash
go run kw-startup.go
```
