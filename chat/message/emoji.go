package message

import(
	"strings"
)

var Emojis = []string{
	":sleeping:", "( ᴗ˳ᴗ) zZ",
	":kiss:",     "( ˘³˘)",
	":relieved:", "( ˘ω˘)",
	":shrug:",    "┐(￣ε￣)┌",
	":unamused:", "( ¬､¬)",
}

func insertEmojis(body string) string {
	r := strings.NewReplacer(Emojis...)
	body = r.Replace(body)
	return body
}
