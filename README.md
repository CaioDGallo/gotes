# Gotes - A Go note taking CLI app for Markdown notes

This CLI app was created while studying the Go language. It solves a problem that I had of context switching for taking meaningful notes. I used to use Obsidian for note taking but now I switched to plain markdown files, and this app aims to facilitate this workflow. It can use the ChatGPT API to create AI-Powered notes from small prompts.

## Usage
- ```gotes new --name "My first note" --subject "My first subject" --content "My first item;My second item"```

The result of this command will be a markdown file with the name "My first note" and the content:

---

# My first subject

## Summary

1. My first item
2. My second item

---

- ```gotes new --name "My first note" --subject "My first subject" --content "My first item;My second item" --ai```

The result of this command will be a markdown file with the name "My first note" and the content:

```Whatever ChatGPT feels like it should create.```

---

## Installation

```shell
go get github.com/caiodgallo/gotes
```

## Configuration

The configuration should be located at ```$HOME/.gotes.yaml```. The default configuration parameters are:

```yaml
rootFolder: "/home/username/your-notes-folder" # Don't use the ~/ home shortcut, use the full path
openAIApiKey: "sk-xxxxx123456xxxx7890"
```

To use the CLI from anywhere you can add the binary to your `/usr/bin` and PATH:

```shell
sudo cp $GOPATH/bin/gotes /usr/local/bin
echo 'export PATH="/usr/bin/gotes:$PATH"' >> "$HOME/.bashrc"
```

---

## Help

```shell
### Usage:
  gotes [flags]
  gotes [command]

### Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  new         Create a new note

### Flags:
      --config string    Config file for Gotes (default is $HOME/.gotes.yaml)
  -h, --help             help for gotes
  -l, --license string   Name of license for the project
      --viper            use Viper for configuration (default true)

Use "gotes [command] --help" for more information about a command.
```
