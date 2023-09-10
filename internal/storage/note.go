package storage

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string
	Content string
}

// Extract all notes from the database
func (storage *Storage) Notes() []*Note {
	var collection []*Note
	storage.database.Find(&collection)
	return collection
}

// Adding a note
func (storage *Storage) AddNote(title, content string) error {
	return storage.database.Create(&Note{
		Title:   title,
		Content: content,
	}).Error
}

// Getting a note
func (storage *Storage) GetNote(ID int) *Note {
	var note Note
	storage.database.First(&note, ID)
	return &note
}

// Deleting a note
func (storage *Storage) DelNote(ID int) {
	storage.database.Delete(&Note{}, ID)
}

// Searches for a note by part of a word in the title
func (storage *Storage) FindNote(part string) []*Note {
	var collection []*Note
	storage.database.Where("lower(title) LIKE ?", "%"+part+"%").Find(&collection)
	return collection
}
