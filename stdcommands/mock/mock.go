package info

import (
	"strings"
	"unicode"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

// mockText capitalizes every second letter and lowercases all other letters
func mockText(input string) string {
	var result strings.Builder
	// Convert the input string to lowercase
	loweredInput := strings.ToLower(input)

	for i, r := range loweredInput {
		// If the index is even, add the character in lowercase
		if i%2 == 0 {
			result.WriteRune(r)
		} else {
			// If the index is odd, capitalize the character
			result.WriteRune(unicode.ToUpper(r))
		}
	}

	return result.String()
}

var Command = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryFun,
	Name:        "Mock",
	Description: "Mocks the text given.",
	Arguments: []*dcmd.ArgDef{
		{Name: "Text", Type: dcmd.String},
	},
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		text := data.Args[0].Str()
		out := mockText(text)
		return out, nil
	},
}
