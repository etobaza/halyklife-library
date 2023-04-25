package models

type Reader struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	FullName      string `json:"full_name"`
	BorrowedBooks []Book `gorm:"many2many:reader_books;" json:"borrowed_books"`
}
