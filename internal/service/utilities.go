package service

import (
	"errors"
	"fmt"
	"gonote/internal/storage"
	"regexp"
	"strings"
)

// Parses the parameter after the command using the Split method
func (service *Service) getParam(command string) (string, error) {
	s := strings.Split(command, " ")

	if len(s) == 1 {
		return s[1], errors.New("parameters are not set")
	}

	return s[1], nil
}

// Parses the parameters after the command using regular expressions
func (service *Service) getParamsByRegex(command string) ([]string, error) {
	re := regexp.MustCompile(`"([^"]+)"`)

	var params []string

	for _, match := range re.FindAllStringSubmatch(command, -1) {
		params = append(params, match[1])
	}

	if len(params) == 0 {
		return params, errors.New("the list of parameters is empty")
	}

	return params, nil
}

// Prints a collection of notes with a pseudo-table
func (service *Service) printTable(collection []*storage.Note) error {
	if len(collection) == 0 {
		return errors.New("the collection is empty")
	}

	for _, note := range collection {
		fmt.Printf("---\nID: %d\tTitle: %s\n---\n", note.ID, note.Title)
	}

	return nil
}
