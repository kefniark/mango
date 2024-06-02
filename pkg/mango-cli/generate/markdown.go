package generate

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/kefniark/mango/pkg/mango-cli/config"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
)

//go:embed templates/*.tmpl
var templates embed.FS

type MarkdownGenerator struct{}

func (generator MarkdownGenerator) Name() string {
	return "Markdown Generator"
}

func (generator MarkdownGenerator) Execute(app string) error {
	tmpl, err := template.ParseFS(templates, "templates/markdown.templ.tmpl")
	if err != nil {
		return err
	}

	folderMango := path.Join(app, "codegen", "views")
	err = os.MkdirAll(folderMango, os.ModePerm)
	if err != nil {
		return err
	}

	config.Logger.Debug().Str("app", app).Str("path", path.Join(folderMango, "markdown.templ")).Msg("Generate Mardown->Templ File")
	f, err := os.Create(path.Join(folderMango, "markdown.templ"))
	if err != nil {
		return err
	}
	defer f.Close()

	cfg := MarkdownConfig{
		Name:    app,
		Entries: []MarkdownConfigEntry{},
	}

	entries, err := os.ReadDir(path.Join(app, "contents"))
	if err != nil {
		return err
	}

	for id, entry := range entries {
		if entry.IsDir() {
			continue
		}

		buf, ctx, err := markdownGenerate(path.Join(app, "contents", entry.Name()))
		if err != nil {
			return err
		}

		metadata := meta.Get(ctx)
		fmt.Println(metadata)

		name := strings.TrimSuffix(entry.Name(), ".md")
		name = strings.TrimSuffix(name, "index")

		cfg.Entries = append(cfg.Entries, MarkdownConfigEntry{
			Title:    entry.Name(),
			URL:      fmt.Sprintf("/%s", name),
			Function: fmt.Sprintf("render_%d", id+1),
			Content:  buf.String(),
		})
	}

	return tmpl.Execute(f, cfg)
}

type MarkdownConfig struct {
	Name    string
	Entries []MarkdownConfigEntry
}

type MarkdownConfigEntry struct {
	Title    string
	URL      string
	Function string
	Content  string
}

func markdownGenerate(src string) (*bytes.Buffer, parser.Context, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			emoji.Emoji,
			meta.Meta,
			&mermaid.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)

	file, err := os.ReadFile(src)
	if err != nil {
		return nil, nil, err
	}

	ctx := parser.NewContext()
	var buf bytes.Buffer
	err = md.Convert(file, &buf, parser.WithContext(ctx))
	return &buf, ctx, err
}
