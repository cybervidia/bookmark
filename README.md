# Bookmark CLI

A simple and lightweight command-line bookmark manager written in Go using Cobra and GORM with SQLite.

## Features

- Add bookmarks by name and URL  
- Add bookmarks by taking URL from system clipboard  
- List all saved bookmarks in a clean table format  
- Retrieve a bookmark's URL and copy it to clipboard  
- Delete bookmarks by name  
- Stores data locally in an SQLite database (`bookmark.db`)  

## Installation

Build the app from source (requires Go 1.18+):

```bash
git clone https://github.com/cybervidia/bookmark.git
cd bookmark
go build -o bookmark main.go
```

Or install directly using:

```bash
go install github.com/cybervidia/bookmark@latest
```

## Usage

Run the `bookmark` command with one of the subcommands:

### Add a bookmark

Add a bookmark with a name and URL:

```bash
bookmark add <name> <url>
```

Add a bookmark with the URL copied from your clipboard:

```bash
bookmark add -c <name>
```

### List all bookmarks

```bash
bookmark list
```

### Get a bookmark URL

Copy the URL of the bookmark to your clipboard:

```bash
bookmark get <name>
```

### Delete a bookmark

Remove a bookmark by name:

```bash
bookmark delete <name>
```

## Example

```bash
bookmark add google https://google.com
bookmark list
bookmark get google
bookmark delete google
```

## Requirements

- Go 1.18 or newer  
- SQLite3 (no additional setup required, included via GORM)  
- Access to system clipboard (handled by `atotto/clipboard` package)  

## License

This project is licensed under the MIT License.

---

Enjoy managing your bookmarks quickly and easily from the terminal!
