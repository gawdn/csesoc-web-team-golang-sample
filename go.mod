module github.com/gawdn/csesoc-web-team-almanac

go 1.12

replace internal/server v0.0.0 => ./internal/server

replace internal/frontmatter v0.0.0 => ./internal/frontmatter

replace pkg/runes v0.0.0 => ./pkg/runes

require (
	github.com/russross/blackfriday v2.0.0+incompatible
	gopkg.in/yaml.v2 v2.2.2
	internal/server v0.0.0
)
