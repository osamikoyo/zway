package models

type Text struct {
	ID    string `gorm:"primarykey"`
	Text  string
	Date  string
	Title string
}
