package models

type Book struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Name     string   `json:"name"`
	Genre    string   `json:"genre"`
	ISBN     string   `json:"isbn"`
	AuthorID uint     `json:"author_id"`
	Author   Author   `json:"-"`
	Readers  []Reader `gorm:"many2many:reader_books;"`
}
