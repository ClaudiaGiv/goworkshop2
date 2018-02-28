package model

import (
	"fmt"
)

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

//Books - the list of available books
var Books []BookDto

// Authors - the list of available authors
var Authors []AuthorDto

type BookList []BookDto
type AuthorsList []AuthorDto

func (b BookDto) String() string{
	return fmt.Sprintf("Book{Title=%s, NoPages=%d, ReleaseDate=%s,Author=%v }",b.Title,b.NoPages,b.ReleaseDate,b.Author)
}
func (b AuthorDto) String() string{
	return fmt.Sprintf("Author{FirstName=%s, LastName=%s, Birthday=%s,Death=%s }",b.FirstName,b.LastName,b.Birthday,b.Death)
}
