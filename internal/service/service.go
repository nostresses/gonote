package service

import (
	"errors"
	"gonote/internal/storage"
	"strings"
)

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{storage: storage}
}

// Executes a service defined through a command from the user
func (service *Service) Execute(command string) error {
	for key, service := range service.availableCommands() {
		if strings.Contains(command, key) {
			service(command)
			return nil
		}
	}
	return errors.New("unknown command")
}

func (service *Service) availableCommands() map[string]Command {
	return map[string]Command{
		"everyone": service.everyone,
		"add":      service.add,
		"get":      service.get,
		"delete":   service.delete,
		"search":   service.search,
	}
}
