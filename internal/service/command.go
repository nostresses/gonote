package service

import (
	"fmt"
	"strconv"
)

type Command func(command string)

func (service *Service) everyone(_ string) {
	collection := service.storage.Notes()

	if err := service.printTable(collection); err != nil {
		fmt.Println(err.Error())
	}
}

func (service *Service) add(command string) {
	params, err := service.getParamsByRegex(command)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(params) < 2 {
		fmt.Println("specify the <title> and <content> parameters")
		return
	}

	if err := service.storage.AddNote(params[0], params[1]); err != nil {
		fmt.Println("couldn't add a note")
		return
	}

	fmt.Println("- Note added.")
}

func (service *Service) get(command string) {
	i, err := service.getParam(command)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		fmt.Println("ID could not be determined")
		return
	}

	note := service.storage.GetNote(id)

	if note.Title == "" || note.Content == "" {
		fmt.Println("Notes not found")
		return
	}

	fmt.Printf("About Note (%d)\n---\nTitle: %s\n---\nContent: %s\n---", note.ID, note.Title, note.Content)
}

func (service *Service) delete(command string) {
	i, err := service.getParam(command)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := strconv.Atoi(i)

	if err != nil {
		fmt.Println("ID could not be determined")
		return
	}

	service.storage.DelNote(id)

	fmt.Println("- Note deleted.")
}

func (service *Service) search(command string) {
	params, err := service.getParamsByRegex(command)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if params[0] == "" {
		fmt.Println("specify the <part> parameter")
		return
	}

	collection := service.storage.FindNote(params[0])

	service.printTable(collection)
}
