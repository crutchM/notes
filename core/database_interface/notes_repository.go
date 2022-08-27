package database_interface

import (
	"github.com/crutchm/notes-core/models"
	"github.com/jmoiron/sqlx"
	"sync"
)

type NotesRepository struct {
	sync.RWMutex
	db *sqlx.DB
}

func (n *NotesRepository) CreateNote(note models.Note) (string, error) {
	var id string
	row := n.db.QueryRow("INSERT INTO notes(id, user, title, body) values ($1,$2,$3,$4) RETURNING id",
		note.Id, note.User, note.Title, note.Body)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (n *NotesRepository) GetNote(id string) (models.Note, error) {
	var note models.Note
	if err := n.db.Get(&note, "SELECT * FROM notes WHERE id=$1", id); err != nil {
		return models.Note{}, err
	}
	return note, nil
}

func (n *NotesRepository) GetAllNotes(userId string) ([]models.Note, error) {
	var notes []models.Note
	if err := n.db.Select(&notes, "SELECT * FROM notes WHERE user_id=$1", userId); err != nil {
		return nil, err
	}
	return notes, nil
}

func NewNotesRepository(db *sqlx.DB) *NotesRepository {
	return &NotesRepository{db: db}
}
