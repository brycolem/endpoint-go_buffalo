package models

type Note struct {
	ID            int    `json:"id" db:"id"`
	NoteText      string `json:"noteText" db:"note_text"`
	ApplicationID int    `json:"applicationId" db:"application_id"`
}
