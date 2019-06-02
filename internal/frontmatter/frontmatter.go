package frontmatter

import (
	"bytes"
	"errors"

	"pkg/runes"

	"gopkg.in/yaml.v2"
)


// ParseFrontmatter parses YAML frontmatter from a byte slice and 
// returns the parsed frontmatter and the remaining content
func ParseFrontmatter(file []byte) (map[interface{}]interface{}, []byte, error) {
	var endOfFrontmatterIndex int

	frontmatterDivider := []rune("---")
	frontmatterContent := make(map[interface{}]interface{})

	dividerIndex := 0
	dividerCount := 0

	fileRunes := bytes.Runes(file)

	for i, fileRune := range fileRunes {

		if dividerCount == 2 {
			endOfFrontmatterIndex = i
			break
		}
		
		if frontmatterDivider[dividerIndex] == fileRune {
			dividerIndex++
		} else {
			dividerIndex = 0
		}

		if dividerIndex == len(frontmatterDivider) {
			dividerIndex = 0
			dividerCount++
		}

	}

	if endOfFrontmatterIndex == 0 {
		return nil, file, errors.New("couldn't find frontmatter")
	}


	frontmatterBytes := runes.ToByteSlice(fileRunes[:endOfFrontmatterIndex])
	yamlErr := yaml.Unmarshal(frontmatterBytes, &frontmatterContent)

	if yamlErr != nil {
		return nil, runes.ToByteSlice(fileRunes[endOfFrontmatterIndex:]), yamlErr
	}

	return frontmatterContent, runes.ToByteSlice(fileRunes[endOfFrontmatterIndex:]), nil
}
