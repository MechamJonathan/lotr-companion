package main

import (
	"fmt"

	"github.com/MechamJonathan/lotr-companion/internal/theoneapi"
	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandGetMovies(cfg *config, args ...string) error {
	movieResp, err := cfg.theoneapiClient.ListMovies()
	if err != nil {
		return err
	}

	printMoviesTable(movieResp)
	return nil
}

func printMoviesTable(mr theoneapi.MovieResponse) {
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(styles.Red))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return styles.HeaderStyle
			case row%2 == 0:
				return styles.EvenRowStyle
			default:
				return styles.OddRowStyle
			}
		}).
		Headers("Movies").Width(72)
	for _, movie := range mr.Docs {
		t.Row(movie.Name)
	}

	fmt.Println("")
	fmt.Println(t)
	fmt.Println("")
}
