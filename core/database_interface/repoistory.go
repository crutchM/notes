package database_interface

import (
	"github.com/crutchm/notes-core/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	CreateUser(user models.User) (string, error)
	GetUser(login, password string) (models.User, error)
}

type NotesRepo interface {
	CreateNote(note models.Note) (string, error)
	GetNote(id string) (models.Note, error)
	GetAllNotes(userId string) ([]models.Note, error)
}

type Repository struct {
	AuthRepo
	NotesRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{AuthRepo: NewAuthRepository(db), NotesRepo: NewNotesRepository(db)}
}
