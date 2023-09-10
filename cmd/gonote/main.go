package main

import (
	"bufio"
	"fmt"
	"gonote/internal/config"
	"gonote/internal/service"
	"gonote/internal/storage"
	"log"
	"os"
)

const help string = "\n\nAvialable commands:\n\n\thelp                        Show available commands\n\teveryone                    Extract all notes from the database\n\tadd \"<title>\" \"<content>\"   Adding a note to the database\n\tget <ID>                    Getting a note from the database by <ID>\n\tdelete <ID>                 Deleting a note from the database by <ID> \n\tsearch \"<part>\"             Search for notes by <part> in the title"

func main() {
	conf, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	storage, err := storage.New(conf.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	service := service.New(storage)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(help)

	for {
		fmt.Print("\n> ")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		text := scanner.Text()

		if text == "quit" {
			break
		}

		if text == "help" {
			fmt.Println(help)
			continue
		}

		if err := service.Execute(scanner.Text()); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
