#Bookmark

## 📌 Bookmark CLI

A simple command-line application to save and manage your bookmarks locally using a SQLite database.

### 🛠️ Usage

```bash
bookmark <name> <url>
```

Saves a URL associated with a given name.

```bash
bookmark -c <name>
bookmark --clip <name>
```

Takes the URL from the clipboard and saves it under the specified name.

```bash
bookmark list
```

Displays a list of all saved bookmarks.
