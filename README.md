# GitIgnore CLI

GitIgnore CLI is a command-line tool that helps you create and manage `.gitignore` files for your projects.

It allows you to easily download updated `.gitignore` files from [github.com/github/gitignore](https://github.com/github/gitignore).

## Installation

Install the tool using `go install`:

```bash
go install github.com/cassiofariasmachado/gitignore-cli@latest
```

> ℹ️ Make sure your Go bin directory (usually `$GOPATH/bin`) is in your PATH environment variable.

## Usage

Run the tool with:

```bash
gitignore-cli --name Go --path /path/where/save/gitignore
```

## Arguments

- `--name`: The name of the language or framework you want to create a `.gitignore` file for. It must match the name of the file in the repository. Examples: [Go](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Go.gitignore), [Python](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Python.gitignore), [Java](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Java.gitignore), etc.
  - Required: `No`
  - Default: `""`

- `--debug`: Enables debug mode. This will print additional information to the console, such as the URL being used to download the `.gitignore` file and the path where it will be saved.
  - Required: `No`
  - Default: `false`

- `--path`: The path where the `.gitignore` file will be saved. If not provided, the file will be saved in the current directory with the default name.
  - Required: `No`
  - Default: `.gitignore`
