local sct = 0
local p = 0
local t = 0
local str = ""

function SelfClosingTagToken(token)
	sct = sct + 1
end

function StartTagToken(token) 
	p = p + 1
end

function EndTagToken(token) 
	p = p - 1
end

function TextToken(token) 
	t = t + 1
end

function Output(str)
	str = str
end

raccoon.Html(raccoon.Token.SelfClosingTag, SelfClosingTagToken)
raccoon.Html(raccoon.Token.StartTag, StartTagToken)
raccoon.Html(raccoon.Token.EndTag, EndTagToken)
raccoon.Html(raccoon.Token.Text, TextToken)
raccoon.Output(Output)
raccoon.Parse()
