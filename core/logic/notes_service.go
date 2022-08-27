package logic

import (
	"github.com/crutchm/notes-core/database_interface"
	"github.com/crutchm/notes-core/models"
)

type NotesService struct {
	repo database_interface.NotesRepo
}

func (n NotesService) CreateNote(note models.InputNote) (string, error) {
	return n.repo.CreateNote(note.MapInputToNote())
}

func (n NotesService) GetNote(id string) (models.Note, error) {
	return n.repo.GetNote(id)
}

func (n NotesService) GetAllNotes(userId string) ([]models.Note, error) {
	return n.repo.GetAllNotes(userId)
}

func NewNotesService(repo database_interface.NotesRepo) *NotesService {
	return &NotesService{repo: repo}
}
