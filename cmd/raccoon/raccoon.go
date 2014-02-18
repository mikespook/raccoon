package main

import (
	"flag"
	"fmt"
	"github.com/mikespook/raccoon"
	"os"
)

var (
	url      string
	script   string
)

func init() {
	flag.StringVar(&url, "url", "", "URL to fatch")
	flag.StringVar(&script, "script", "", "Lua script for parsing")
	flag.Parse()
}

func main() {
	if url == "" || script == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	r := raccoon.New(url)
	l := raccoon.LuaWrap(r)
	defer l.Close()
	if err := l.DoFile(script); err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}
}
