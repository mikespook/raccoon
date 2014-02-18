package raccoon

import (
	"testing"
	"io/ioutil"
)

var (
	luafile = "examples/foobar.lua"
)

func TestLuaDoFile(t *testing.T) {
	r := New("http://www.example.com/")
	l := LuaWrap(r)
	if err := l.DoFile(luafile); err != nil {
		t.Error(err)
		return
	}
}

func TestLuaDoString(t *testing.T) {
	r := New("http://www.example.com/")
	l := LuaWrap(r)
	b, err := ioutil.ReadFile(luafile)
	if err != nil {
		t.Error(err)
		return
	}
	if err := l.DoString(string(b)); err != nil {
		t.Error(err)
		return
	}
}
