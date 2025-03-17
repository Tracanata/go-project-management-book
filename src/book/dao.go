package book

type Book struct {
	Code_Book     string `json:"codeBook" gorm:"column:code_book"`
	Title_Book    string `json:"titleBook" gorm:"column:title_book"`
	Author        string `json:"author" gorm:"column:author"`
	Realease_Year int    `json:"year" gorm:"column:realease_year"`
}

func (Book) TableName() string {
	return "book"
}
