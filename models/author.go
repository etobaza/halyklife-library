package models

type Author struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	FullName       string `json:"full_name"`
	Nickname       string `json:"nickname"`
	Specialization string `json:"specialization"`
	Books          []Book `gorm:"ForeignKey:AuthorID"`
}
