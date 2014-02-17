package raccoon

import (
	"code.google.com/p/go.net/html"
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
	"net/http"
)

const (
	module = "raccoon"
)

type Raccoon struct {
	Url, Script string
	luaState    *lua.State
	parseMap    map[html.TokenType]*luar.LuaObject
	tokenizer   *html.Tokenizer
}

func New(url, script string) (raccoon *Raccoon, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	raccoon = &Raccoon{
		Url:       url,
		Script:    script,
		luaState:  luar.Init(),
		parseMap:  make(map[html.TokenType]*luar.LuaObject),
		tokenizer: html.NewTokenizer(resp.Body),
	}
	m := luar.Map{
		"URL":       raccoon.Url,
		"Script":    raccoon.Script,
		"Tokenizer": raccoon.tokenizer,
		"IsError":   isError,
		"Token": map[string]html.TokenType{
			"SelfClosingTag": html.SelfClosingTagToken,
			"StartTag":       html.StartTagToken,
			"EndTag":         html.EndTagToken,
			"Text":           html.TextToken,
		},
		"Parse": raccoon.parse,
		"Bind":  raccoon.bind,
	}
	luar.Register(raccoon.luaState, module, m)
	return
}

func (raccoon *Raccoon) bind(t html.TokenType, f interface{}) {
	if obj, ok := f.(*luar.LuaObject); ok && obj.Type == "function" {
		raccoon.parseMap[t] = obj
	} else {
		raccoon.luaState.RaiseError("Not function")
	}
}

func (raccoon *Raccoon) parse() {
	for {
		tokenType := raccoon.tokenizer.Next()
		if tokenType == html.ErrorToken {
			return
		}
		if obj, ok := raccoon.parseMap[tokenType]; ok {
			token := raccoon.tokenizer.Token()
			m := luar.Map{
				"Type":     token.Type,
				"DataAtom": token.DataAtom.String(),
				"Data":     token.String(),
				"Attr":     luar.NewLuaObjectFromValue(raccoon.luaState, token.Attr),
			}
			if _, err := obj.Call(m); err != nil {
				raccoon.luaState.RaiseError(token.Type.String() + ": " + err.Error())
			}
		}
	}
}

func (raccoon *Raccoon) Parse() error {
	return raccoon.luaState.DoFile(raccoon.Script)
}

func (raccoon *Raccoon) Close() {
	raccoon.luaState.Close()
}

func (raccoon *Raccoon) LuaState() *lua.State {
	return raccoon.luaState
}

func isError(tokenType html.TokenType) bool {
	return tokenType == html.ErrorToken
}
