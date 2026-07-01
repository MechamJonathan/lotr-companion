package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MechamJonathan/lotr-companion/internal/theoneapi"
	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/muesli/reflow/wordwrap"
)

var quotesHelpMessage = styles.ErrorMessage.SetString("enter 'quotes' to go forward a page. 'quotesb' to go back a page.\n or enter 'quotes <character name>' to view another character's quotes.").Render()

func commandQuotesf(cfg *config, args ...string) error {
	inputName := strings.Join(args, " ")

	if len(args) < 1 && cfg.currentCharacterName == "" {
		cmdUsage := "usage: quotes <character name>\n"
		cmdOptions := [][]string{
			{"character name", "retrieve quotes from specific character"},
		}
		PrintUsageTable(cmdUsage, cmdOptions)
		return nil

	} else if cfg.currentCharacterName == "" {
		cfg.currentCharacterName = inputName
	} else if cfg.currentCharacterName != "" && len(args) >= 1 {
		cfg.currentCharacterName = inputName
		cfg.currentQuotePage = 0
	}

	cfg.currentQuotePage += 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentCharacterName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	printQuotesTable(cfg.currentQuotePage, quotesResp.Docs, cfg.currentCharacterName)

	return nil
}

func commandQuotesb(cfg *config, args ...string) error {
	if cfg.currentQuotePage <= 1 {
		cfg.currentQuotePage -= 1
		return errors.New("you're on the first page of quotes")
	}

	cfg.currentQuotePage -= 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentCharacterName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	printQuotesTable(cfg.currentQuotePage, quotesResp.Docs, cfg.currentCharacterName)
	return nil
}

func printQuotesTable(page int, quotes []theoneapi.Quote, characterName string) {
	headerString := fmt.Sprintf("Page: %v of %s Quotes", page, characterName)
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(styles.Red))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return styles.HeaderStyle
			case row%2 == 0:
				return styles.QuoteStyle
			default:
				return styles.OddQuoteStyle
			}
		}).
		Headers(headerString).Width(72)
	for _, quote := range quotes {
		wrappedText := wordwrap.String(quote.Dialog, 60)
		t.Row(wrappedText)
	}

	fmt.Println(t)
	fmt.Println(quotesHelpMessage)
}
