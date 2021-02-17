package message

import(
	"strings"
)

var EmojiMap = map[string]string{
	"( ᴗ˳ᴗ) zZ": "sleeping",
	"( ˘³˘)":    "kiss",
	"( ˘ω˘)":    "relieved",
	"┐(￣ε￣)┌": "shrug",
	"( ¬､¬)":    "unamused",
}

func insertEmojis(body string) string {
	for emoji, name := range(EmojiMap) {
		body = strings.ReplaceAll(body, ":" + name + ":", emoji)
	}
	return body
}
