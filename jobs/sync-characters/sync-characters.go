package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing characters with Kitsu DB")
	defer arn.Node.Close()

	kitsuCharacters := kitsu.StreamCharacters()

	for kitsuCharacter := range kitsuCharacters {
		character := &arn.Character{
			ID:          kitsuCharacter.ID,
			Name:        kitsuCharacter.Attributes.Name,
			Image:       kitsu.FixImageURL(kitsuCharacter.Attributes.Image.Original),
			Description: kitsuCharacter.Attributes.Description,
			Attributes:  []*arn.CharacterAttribute{},
		}

		// We use markdown, so replace <br/> with two line breaks.
		character.Description = strings.Replace(character.Description, "<br/>", "\n\n", -1)

		// Parse attributes like these:
		// - Position: Club Manager
		// - Height: 162 cm (5' 4")
		// - Weight: 48 kg (106 lb)
		// - Birthday: November 24
		// - Hair color: Brown
		// - Eyes: Blue (anime), Green (manga)

		lines := strings.Split(character.Description, "\n\n")
		finalLines := make([]string, 0, len(lines))

		for _, line := range lines {
			originalLine := line

			if strings.HasPrefix(line, "(") {
				line = strings.TrimPrefix(line, "(")
				line = strings.TrimSuffix(line, ")")
			}

			line = strings.TrimSpace(line)

			colonPos := strings.Index(line, ":")

			if colonPos == -1 || colonPos < 2 || colonPos > 25 {
				finalLines = append(finalLines, originalLine)
				continue
			}

			key := line[:colonPos]
			value := line[colonPos+1:]

			value = strings.TrimSpace(value)

			if key == "source" {
				key = "Source"
			}

			character.Attributes = append(character.Attributes, &arn.CharacterAttribute{
				Name:  key,
				Value: value,
			})

			fmt.Println(color.CyanString(key), color.YellowString(value))
		}

		character.Description = strings.Join(finalLines, "\n\n")
		character.Description = strings.Trim(character.Description, "\n\n")
		character.Save()

		fmt.Printf("%s %s %s\n", color.GreenString("✔"), character.ID, character.Name)
	}

	color.Green("Finished.")
}
