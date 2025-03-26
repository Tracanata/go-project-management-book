package book

type RespDataBook struct {
	Code_Book  string `json:"codeBook"`
	Title_Book string `json:"titleBook"`
	Author     string `json:"author"`
	Stock      int    `json:"stock"`
}
