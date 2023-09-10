# gonote

**gonote** is a simple manager of your notes.


### Installation

1. Install and unzip the archive
1.2 or use `git clone <repo> <directory>`
2. Go to the directory with the project (`cd "gonote"`) and run the command:
```sh
go run cmd/gonote/main.go
```

### Usage

**Avialable commands**

```
help                        Show available commands
everyone                    Extract all notes from the database
add "<title>" "<content>"   Adding a note to the database
get <ID>                    Getting a note from the database by <ID>
delete <ID>                 Deleting a note from the database by <ID>
search "<part>"             Search for notes by <part> in the title
```

**Examples**

```sh
> everyone
# ...

> add "Title" "Content"
# ...

> get 5
# ...

> delete 1
# ...

> search "part of the title"
# ...
```