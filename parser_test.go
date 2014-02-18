package raccoon

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"testing"
)

func TestHtmlParse(t *testing.T) {
	r := New("http://www.example.com/")
	r.Html(html.SelfClosingTagToken, func(token html.Token) error {
		if token.DataAtom.String() != "meta" {
			return fmt.Errorf("meta is expected, got %s", token.DataAtom.String())
		}
		return nil
	})
	if err := r.Parse(); err != nil {
		t.Error(err)
	}
}

func TestHtmlParseError(t *testing.T) {
	r := New("http://www.example.com/")
	expErr := fmt.Errorf("Expected error")
	r.Html(html.SelfClosingTagToken, func(token html.Token) error {
		return expErr
	})
	if err := r.Parse(); err != expErr {
		t.Error("Parse should return an error")
	}
}

func TestOutputTo(t *testing.T) {
	r := New("http://www.example.com/")
	strbuf := bytes.NewBufferString("")
	r.OutputTo(strbuf)
	if err := r.Parse(); err != nil {
		t.Error(err)
		return
	}
	if strbuf.Len() == 0 {
		t.Errorf("Empty buffer")
	}
}

func TestHtmlParseAndOutputTo(t *testing.T) {
	r := New("http://www.example.com/")
	strbuf := bytes.NewBufferString("")
	r.OutputTo(strbuf)
	r.Html(html.SelfClosingTagToken, func(token html.Token) error {
		if token.DataAtom.String() != "meta" {
			return fmt.Errorf("meta is expected, got %s", token.DataAtom.String())
		}
		return nil
	})
	if err := r.Parse(); err != nil {
		t.Error(err)
		return
	}
	if strbuf.Len() == 0 {
		t.Errorf("Empty buffer")
	}
}
