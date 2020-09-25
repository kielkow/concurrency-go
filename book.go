package main

import "fmt"

// Book struct
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Title 1",
		Author:        "Author 1",
		YearPublished: 2000,
	},
	Book{
		ID:            2,
		Title:         "Title 2",
		Author:        "Author 2",
		YearPublished: 2000,
	},
	Book{
		ID:            3,
		Title:         "Title 3",
		Author:        "Author 3",
		YearPublished: 2000,
	},
}
