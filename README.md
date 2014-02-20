Raccoon
=======

[![Build Status][travis-img]][travis]

Raccoon is a simple web-spider framework driven by Golang and Lua.

Install
=======

Install the package:

> $ go get github.com/mikespook/raccoon

Install the CLI command:

> $go get github.com/mikespook/raccoon/cmd/raccoon

Usage
=====

Raccoon's parser can be used for parsing html in Golang application directly:

```go
p := raccoon.New("http://www.example.com/")
p.Html(html.SelfClosingTagToken, func(token html.Token) error {
	if token.DataAtom.String() != "meta" {
		return fmt.Errorf("meta is expected, got %s", token.DataAtom.String())
	}
	return nil
})
if err := r.Parse(); err != nil {
	// handling error
}
```

Or using lua scripts:

```go
p := raccoon.New("http://www.example.com/")
l := raccoon.LuaWrap(p)
if err := l.DoFile("foobar.lua"); err != nil {
	// handling error
}
```

```lua
function SelfClosingTagToken(token)
	if token.DataAtom ~= "meta" then
		error("meta is expected, got " .. token.DataAtom)
	end
end

raccoon.Html(raccoon.Token.SelfClosingTag, SelfClosingTagToken)
raccoon.Parse()
```

Contributors
============

(_Alphabetic order_)
 
 * [Xing Xing][blog] <mikespook@gmail.com> [@Twitter][twitter]

Open Source - MIT Software License
==================================

See LICENSE.

 [travis-img]: https://travis-ci.org/mikespook/raccoon.png?branch=master
 [travis]: https://travis-ci.org/mikespook/raccoon
 [blog]: http://mikespook.com
 [twitter]: http://twitter.com/mikespook
