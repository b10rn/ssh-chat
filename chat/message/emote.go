package message

import(
	"strings"
)

var Emotes = []string{
	":sleeping:", "( ᴗ˳ᴗ) zZ",
	":kiss:",     "( ˘³˘)",
	":relieved:", "( ˘ω˘)",
	":shrug:",    "┐(￣ε￣)┌",
	":unamused:", "( ¬､¬)",
}

func insertEmotes(body string) string {
	r := strings.NewReplacer(Emotes...)
	body = r.Replace(body)
	return body
}
