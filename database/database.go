package database

import "errors"

// Define the Author struct and create a sample dataset

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Pseudonym   string `json:"pseudonym"`
	Description string `json:"description"`
}

var data = [5]Author{
	{ID: "4a7a46a8-e704-41d0-be93-6ec395d32254", Name: "John Doe", Pseudonym: "Jane Doe", Description: "Aspiring young author."},
	{ID: "4a7a46a8-e704-41d0-be93-6ec395d32255", Name: "H. G. Wells", Pseudonym: "n/a", Description: "n/a"},
	{ID: "4a7a46a8-e704-41d0-be93-6ec395d32256", Name: "Andy Weir", Pseudonym: "n/a", Description: "n/a"},
	{ID: "4a7a46a8-e704-41d0-be93-6ec395d32257", Name: "Tillie Walden", Pseudonym: "n/a", Description: "Comic artist."},
	{ID: "4a7a46a8-e704-41d0-be93-6ec395d32258", Name: "J. R. R. Tolkien", Pseudonym: "n/a", Description: "Lord of the rings."},
}

/*
	GetAuthor gets an author given a string ID
*/
func GetAuthor(id string) (*Author, error) {
	i := -1
	for index, el := range data {
		if el.ID == id {
			i = index
		}
	}

	if i == -1 {
		return nil, errors.New("could not find author for given id")
	}
	return &data[i], nil
}

/*
	GetAuthors gets the entire list of authors
*/
func GetAuthors() [5]Author {
	return data
}
