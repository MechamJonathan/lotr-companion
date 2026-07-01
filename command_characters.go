package main

import (
	"fmt"

	"github.com/MechamJonathan/lotr-companion/internal/theoneapi"
	"github.com/MechamJonathan/lotr-companion/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandGetCharacters(cfg *config, args ...string) error {
	if len(args) < 1 {
		cmdUsage := "usage: characters <group>"
		cmdOptions := [][]string{
			{"all", "All characters"},
			{"fellowship", "Members of the Fellowship of the Ring"},
			{"hobbits", "Popular Hobbits"},
			{"men", "Popular Men (and Women) of Middle-earth "},
			{"elves", "Popular Elves"},
			{"dwarves", "Popular Dwarves"},
			{"orcs", "Popular Orcs and Goblins"},
			{"wizards", "The Istari (wizards)"},
			{"creatures", "Other creatures and beasts"},
		}

		PrintUsageTable(cmdUsage, cmdOptions)
		return nil
	}

	charResp, err := cfg.theoneapiClient.ListCharacters()
	if err != nil {
		return err
	}

	switch args[0] {
	case "all":
		printAllCharacters(charResp.Docs)
	case "fellowship":
		getFellowshipMembers(charResp.Docs)
	case "hobbits":
		getHobbitMembers(charResp.Docs)
	case "men":
		getMenOfMiddleEarth(charResp.Docs)
	case "elves":
		getElves(charResp.Docs)
	case "dwarves":
		getDwarves(charResp.Docs)
	case "orcs":
		getOrcs(charResp.Docs)
	case "wizards":
		getWizards(charResp.Docs)
	case "creatures":
		getCreatures(charResp.Docs)
	default:
		return fmt.Errorf("invalid option: '%s'", args[0])
	}

	return nil
}

func printAllCharacters(characters []theoneapi.Character) {
	var allCharacters []string
	for _, character := range characters {
		allCharacters = append(allCharacters, character.Name)
	}
	printGroupMembersTable("All Characters", allCharacters, characters)
}

func getFellowshipMembers(characters []theoneapi.Character) {
	fellowshipMembers := []string{
		"Frodo Baggins", "Samwise Gamgee", "Gandalf", "Aragorn II Elessar", "Legolas",
		"Gimli", "Boromir", "Meriadoc Brandybuck", "Peregrin Took",
	}
	printGroupMembersTable("Fellowship Members", fellowshipMembers, characters)
}

func getHobbitMembers(characters []theoneapi.Character) {
	hobbitMembers := []string{
		"Frodo Baggins", "Samwise Gamgee", "Meriadoc Brandybuck", "Peregrin Took", "Bilbo Baggins",
	}
	printGroupMembersTable("Hobbits", hobbitMembers, characters)
}

func getMenOfMiddleEarth(characters []theoneapi.Character) {
	menOfMiddleEarth := []string{
		"Aragorn II Elessar", "Boromir", "Faramir", "Théoden", "Éomer", "Éowyn", "Denethor", "Bard",
		"Gríma Wormtongue", "Denethor II",
	}
	printGroupMembersTable("Men of Middle Earth", menOfMiddleEarth, characters)
}

func getElves(characters []theoneapi.Character) {
	elves := []string{
		"Legolas", "Elrond", "Galadriel", "Arwen", "Thranduil", "Glorfindel", "Haldir", "Celeborn",
	}
	printGroupMembersTable("Elves", elves, characters)
}

func getDwarves(characters []theoneapi.Character) {
	dwarves := []string{
		"Gimli", "Thorin II Oakenshield", "Balin", "Dwalin", "Bofur", "Durin", "Dáin II Ironfoot",
		"Fíli and Kíli", "Óin", "Glóin", "Bifur", "Bombur", "Dori", "Nori", "Ori",
	}
	printGroupMembersTable("Dwarves", dwarves, characters)
}

func getWizards(characters []theoneapi.Character) {
	wizards := []string{
		"Gandalf", "Saruman", "Sauron", "Radagast", "Alatar", "Pallando",
	}
	printGroupMembersTable("Wizards", wizards, characters)
}

func getOrcs(characters []theoneapi.Character) {
	orcs := []string{
		"Azog", "Bolg", "Gothmog", "Uglúk", "Grishnákh", "Shagrat", "Gorbag", "Snaga",
	}
	printGroupMembersTable("Orcs", orcs, characters)
}

func getCreatures(characters []theoneapi.Character) {
	creatures := []string{
		"Gollum", "Smaug", "Shelob", "Treebeard", "Watcher in the Water", "Gwaihir", "Durin's Bane", "Witch-King of Angmar",
		"Khamûl",
	}
	printGroupMembersTable("Creatures", creatures, characters)

}

func printGroupMembersTable(title string, groupMembers []string, characters []theoneapi.Character) {
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
		Headers(title).
		Width(72)

	for _, character := range characters {
		for _, member := range groupMembers {
			if character.Name == member {
				t.Row(character.Name)
			}
		}
	}

	fmt.Println("")
	fmt.Println(t)
	fmt.Println("")
}
