package main

import (
	"fmt"
	"sort"

	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println(styles.Title.Render("Welcome to Palantir!"))
	fmt.Println(lipgloss.NewStyle().SetString(",---.\n<(  0  )>\n`---'").Align(lipgloss.Center).Width(70).Foreground(styles.Orange))
	fmt.Println(styles.SubHeader.Render("━━━ ✧༚ The unofficial LOTR companion app. ✧༚ ━━━ \n Made possible by The-One-Api."))

	commands := getCommands()

	var commandNames []string
	for name := range commands {
		commandNames = append(commandNames, name)
	}

	sort.Strings(commandNames)

	printCommandHelpTable(commandNames, commands)
	return nil
}

func printCommandHelpTable(commandNames []string, commands map[string]cliCommand) {
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
		Headers("Command", "Description")

	for _, name := range commandNames {
		cmd := commands[name]
		t.Row(cmd.name, cmd.description)
	}
	fmt.Println(t)
	fmt.Println("")
}
