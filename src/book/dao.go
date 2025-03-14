package book

type Book struct {
	Id            int    `json:"id"`
	Title_Book    string `json:"titleBook"`
	Author        string `json:"author"`
	Realease_Year int    `json:"year"`
}

func (Book) TableName() string {
	return "book"
}
