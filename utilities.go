package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"

	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

func ClearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to clear screen")
	}
	return nil
}

func MoveCursorToBottom() {
	th := GetTerminalHeight()
	fmt.Printf("\033[%d;1H", th)
}

func GetTerminalHeight() int {
	height, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 20
	}
	return height
}

func getRandomQuote(quotes []string) string {
	var index uint32
	err := binary.Read(rand.Reader, binary.BigEndian, &index)
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return ""
	}

	return quotes[int(index)%len(quotes)]
}

func PrintUsageTable(cmdUsage string, options [][]string) {
	if err := ClearScreen(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(styles.Title.Render(cmdUsage))
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
		Headers("Options", "Description").
		Width(72)

	t.Rows(options...)
	fmt.Println(t)
}
