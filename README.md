# GitIgnore CLI

GitIgnore CLI is a command-line tool that helps you create and manage `.gitignore` files for your projects.

It allows you to easily download updated `.gitignore` files from [github.com/github/gitignore](https://github.com/github/gitignore).

## Usage

To use GitIgnore CLI, simply run the following commands in your terminal:

1. Install the tool using `go install`:
    ```bash
    go install github.com/cassiofariasmachado/gitignore-cli
    ```

2. Run the tool with the desired command:
    ```bash
    gitignore-cli --name Go
    ```

## Arguments

- `-name`: The name of the language or framework you want to create a `.gitignore` file for. It must match the name of the file in the repository. Examples: [Go](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Go.gitignore), [Python](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Python.gitignore), [Java](https://raw.githubusercontent.com/github/gitignore/refs/heads/main/Java.gitignore), etc.
  - Required: `No`
  - Default: `""`

- `-debug`: Enables debug mode. This will print additional information to the console, such as the URL being used to download the `.gitignore` file and the path where it will be saved.
  - Required: `No`
  - Default: `false`

- `-path`: The path where the `.gitignore` file will be saved. If not provided, the file will be saved in the current directory with the default name.
  - Required: `No`
  - Default: `.gitignore`
