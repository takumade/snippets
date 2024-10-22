package main

import "snippetbox.takucoder.dev/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}