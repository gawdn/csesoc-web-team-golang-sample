module github.com/gawdn/csesoc-web-team-almanac

go 1.12

replace internal/server v0.0.0 => ./internal/server

replace internal/frontmatter v0.0.0 => ./internal/frontmatter

replace pkg/runes v0.0.0 => ./pkg/runes

require (
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/github_flavored_markdown v0.0.0-20181002035957-2122de532470 // indirect
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20181215221002-9d8641ddf2e1 // indirect
	github.com/shurcooL/octicon v0.0.0-20181222203144-9ff1a4cf27f4 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	gopkg.in/yaml.v2 v2.2.2
	internal/server v0.0.0
)
