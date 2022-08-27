package models

import "github.com/google/uuid"

type Note struct {
	Id    string `json:"id" db:"id"`
	User  string `json:"user_id" db:"user_id"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body"`
}

type InputNote struct {
	User  string `json:"user_id" db:"user_id"`
	Title string `json:"title" db:"title"`
	Body  string `json:"body" db:"body"`
}

func (s *InputNote) MapInputToNote() Note {
	newNote := Note{
		Id:    uuid.NewString()[:8],
		User:  s.User,
		Title: s.Title,
		Body:  s.Body,
	}
	return newNote
}
