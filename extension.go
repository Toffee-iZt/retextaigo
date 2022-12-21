package retextaigo

import "github.com/karalef/retextaigo/api"

// Extended is extended response.
type Extended []any

// Complete returns complete text.
// It uses first string from each extension variant.
func (e Extended) Complete() string {
	l := 0
	for _, v := range e {
		switch a := v.(type) {
		case string:
			l += len(a)
		case []string:
			l += len(a[0])
		}
	}
	buf := make([]byte, 0, l+len(e)-1)
	for i, v := range e {
		switch a := v.(type) {
		case string:
			buf = append(buf, a...)
		case []string:
			buf = append(buf, a[0]...)
		}
		if i < len(e)-1 {
			buf = append(buf, ' ')
		}
	}
	return string(buf)
}

// Extension generates extended text for source.
func (c *Client) Extension(source string, lang string) (*Task[[]string], error) {
	var add map[string]any
	if lang != "" {
		add = map[string]any{"lang": lang}
	}
	return queue[[]string](c, api.TaskExtension, source, add)
}
