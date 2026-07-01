package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandGetDetails(cfg *config, args ...string) error {
	if len(args) < 1 {
		cmdUsage := "usage: details <option>"
		cmdOptions := [][]string{
			{"character name", "retrieve a specific character's details"},
			{"movie", "retreive a movie's details or a movie series' details"},
			{"book", "retreive a book's details"},
		}

		PrintUsageTable(cmdUsage, cmdOptions)
		return nil
	}

	inputName := strings.Join(args, " ")

	err := fetchCharacterDetails(cfg, inputName)
	if err == nil {
		return nil
	}

	movieErr := fetchMovieDetails(cfg, inputName)
	_, _, bookErr := fetchBookDetails(cfg, inputName)
	if movieErr == nil || bookErr == nil {
		return nil
	}
	return fmt.Errorf("no details found for: %s", inputName)
}

func fetchMovieDetails(cfg *config, name string) error {
	movieResp, err := cfg.theoneapiClient.GetMovieByName(name)
	if err != nil {
		return err
	}

	runtime := fmt.Sprint(movieResp.RuntimeInMinutes)
	budget := strconv.FormatFloat(movieResp.BudgetInMillions, 'f', -1, 64)
	boxOffice := strconv.FormatFloat(movieResp.BoxOfficeRevenueInMillions, 'f', -1, 64)
	rottenTomatoesScore := strconv.FormatFloat(movieResp.RottenTomatoesScore, 'f', -1, 64)
	awards := fmt.Sprintf("Awards: %d nominations, %d wins", movieResp.AcademyAwardNominations, movieResp.AcademyAwardWins)

	rows := [][]string{
		{"Name", movieResp.Name},
		{"Runtime", runtime + " mins"},
		{"Budget", "$" + budget + "M"},
		{"Box Office", "$" + boxOffice + "M"},
		{"Awards", awards},
		{"Rotten Tomatos", rottenTomatoesScore},
	}

	movieTableName := movieResp.Name + " (movie)"
	bookRows, bookTableName, _ := fetchBookDetails(cfg, name)
	if bookRows != nil {
		printMoviesAndBooksDetailsTables(rows, movieTableName, bookRows, bookTableName)
	} else {
		printMoviesAndBooksDetailsTables(rows, movieTableName, nil, "")
	}
	return nil
}

func fetchBookDetails(cfg *config, name string) ([][]string, string, error) {
	bookResp, err := cfg.theoneapiClient.GetBookByName(name)
	if err != nil {
		return nil, "", err
	}

	rows := [][]string{
		{"Name", bookResp.Name},
		{"", "(No additional details availble currently)"},
	}

	bookTableName := bookResp.Name + " (book)"
	return rows, bookTableName, err
}

func fetchCharacterDetails(cfg *config, name string) error {
	charResp, err := cfg.theoneapiClient.GetCharacterByName(name)
	if err != nil {
		return err
	}

	rows := [][]string{
		{"Name", charResp.Name},
		{"WikiURL", charResp.WikiURL},
		{"Race", charResp.Race},
		{"Birth", charResp.Birth},
		{"Gender", charResp.Gender},
		{"Death", charResp.Death},
		{"Hair", charResp.Hair},
		{"Height", charResp.Height},
		{"Realm", charResp.Realm},
		{"Spouse", charResp.Spouse},
	}

	printDetailsTable(rows, charResp.Name)
	return nil
}

func printMoviesAndBooksDetailsTables(movieRows [][]string, movieTableTitle string, bookRows [][]string, bookTableTitle string) {
	if movieRows != nil {
		printDetailsTable(movieRows, movieTableTitle)
	}
	if bookRows != nil {
		printDetailsTable(bookRows, bookTableTitle)
	}
}

func printDetailsTable(rows [][]string, name string) {
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
		Headers("", name).
		Width(72).
		Rows(rows...)

	fmt.Println("")
	fmt.Println(t)
	fmt.Println("")
}
