package raccoon

import (
	"code.google.com/p/go.net/html"
	"io"
	"net/http"
	"strings"
)

const (
	module = "raccoon"
)

type TokenHandler func(html.Token) error

type parser struct {
	url     string
	htmlMap map[html.TokenType]TokenHandler
	writer  io.Writer
}

func New(url string) (p *parser) {
	return &parser{url, make(map[html.TokenType]TokenHandler), nil}
}

func (p *parser) Html(t html.TokenType, f TokenHandler) {
	p.htmlMap[t] = f
}

func (p *parser) OutputTo(w io.Writer) {
	p.writer = w
}

func (p *parser) Parse() error {
	resp, err := http.Get(p.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if p.writer != nil {
		if _, err := io.Copy(p.writer, resp.Body); err != nil {
			return err
		}
	}
	if strings.Contains(resp.Header.Get("Content-Type"), "html") {
		tokenizer := html.NewTokenizer(resp.Body)
		if err := p.html(tokenizer); err != nil {
			return err
		}
	}
	return nil
}

func (p *parser) html(tokenizer *html.Tokenizer) error {
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if err := tokenizer.Err(); err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		if f, ok := p.htmlMap[tokenType]; ok {
			token := tokenizer.Token()
			if err := f(token); err != nil {
				return err
			}
		}
	}
	return nil
}
