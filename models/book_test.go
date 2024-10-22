package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookStruct(t *testing.T) {
	book := Book{
		Title:       "Test Book",
		Author:      "Test Author",
		Description: "Test Description",
		PublishedAt: "2024-10-16",
	}

	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, "Test Author", book.Author)
	assert.Equal(t, "Test Description", book.Description)
	assert.Equal(t, "2024-10-16", book.PublishedAt)
}
