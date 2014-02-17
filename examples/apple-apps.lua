local paidApps = {}
local capture = 0
local tmp = {}
local enter = ""

function SelfClosingTagToken(token)
	if capture > 0 then
		print(token.Type.String(), token.DataAtom, token.Data, #token.Attr)
	end
end

function StartTagToken(token) 
	if token.DataAtom == "section" then
		capture = 1
	end
	if capture > 0 then
		if token.DataAtom ~= "img" then
			capture = capture + 1
		end
	end

	if capture == 4 then
		tmp = {}
	end

	if capture == 5 then
		enter = token.DataAtom
		if enter == "a" then
			tmp.Url = token.Attr[1].Val
		end
		if token.DataAtom == "img" then
			tmp.Img = token.Attr[1].Val
		end
	end
end

function EndTagToken(token) 
	if token.DataAtom == "section" then
		capture = 0
	end

	if capture > 0 then 
		if token.DataAtom ~= "img" then
			capture = capture - 1
		end
	end

	if capture == 3 then
		paidApps[tmp.Index] = tmp
	end
end

function TextToken(token) 
	if capture == 5 then
		if enter == "strong" then
			tmp.Index = tonumber(token.Data)
		end
	end
	if capture == 6 then
		if enter == "h3" then
			tmp.Title = token.Data			
		end
		if enter == "h4" then
			tmp.Category = token.Data
		end
	end
end

local err

raccoon.Bind(raccoon.Token.SelfClosingTag, SelfClosingTagToken)
raccoon.Bind(raccoon.Token.StartTag, StartTagToken)
raccoon.Bind(raccoon.Token.EndTag, EndTagToken)
raccoon.Bind(raccoon.Token.Text, TextToken)
raccoon.Parse()

json = require("json")
print(json.encode(paidApps))
