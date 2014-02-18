package raccoon

import (
	"code.google.com/p/go.net/html"
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
)

const (
	module = "raccoon"
)

type luaWriter struct {
	state *lua.State
	obj   *luar.LuaObject
}

func (w *luaWriter) Write(p []byte) (n int, err error) {
	if _, err := w.obj.Call(string(p)); err != nil {
		w.state.RaiseError(err.Error())
		return 0, err
	}
	return len(p), nil
}

type luaWrap struct {
	state *lua.State
	p     *parser
}

func LuaWrap(p *parser) *luaWrap {
	l := &luaWrap{luar.Init(), p}
	m := luar.Map{
		"URL": l.p.url,
		"Token": map[string]html.TokenType{
			"SelfClosingTag": html.SelfClosingTagToken,
			"StartTag":       html.StartTagToken,
			"EndTag":         html.EndTagToken,
			"Text":           html.TextToken,
		},
		"Output": func(f interface{}) {
			if obj, ok := f.(*luar.LuaObject); ok && obj.Type == "function" {
				l.p.OutputTo(&luaWriter{l.state, obj})
			} else {
				l.state.RaiseError("Arg #2 is not a function.")
			}
		},
		"Parse": l.p.Parse,
		"Html": func(t html.TokenType, f interface{}) {
			if obj, ok := f.(*luar.LuaObject); ok && obj.Type == "function" {
				l.p.Html(t, func(token html.Token) error {
					m := luar.Map{
						"Type":     token.Type,
						"DataAtom": token.DataAtom.String(),
						"Data":     token.String(),
						"Attr":     luar.NewLuaObjectFromValue(l.state, token.Attr),
					}
					if _, err := obj.Call(m); err != nil {
						l.state.RaiseError(token.Type.String() + ": " + err.Error())
					}
					return nil
				})
			} else {
				l.state.RaiseError("Arg #2 is not a function.")
			}
		},
	}
	luar.Register(l.state, module, m)
	return l
}

func (l *luaWrap) DoFile(filename string) error {
	m := luar.Map{
		"FileName": filename,
	}
	luar.Register(l.state, module, m)
	if err := l.state.DoFile(filename); err != nil {
		return err
	}
	return nil
}

func (l *luaWrap) DoString(str string) error {
	if err := l.state.DoString(str); err != nil {
		return err
	}
	return nil
}

func (l *luaWrap) Close() {
	l.state.Close()
}
