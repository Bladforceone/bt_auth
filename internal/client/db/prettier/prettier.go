package prettier

import (
	"fmt"
	"strings"
)

const (
	PlaceHolderDollar = "$"
)

func Pretty(query string, placeholder string, args ...any) string {
	for i, arg := range args {
		var value string
		switch v := arg.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}

		query = strings.Replace(query, fmt.Sprintf("%s%d", placeholder, i+1), value, -1)
	}

	query = strings.Replace(query, "\t", "", -1)
	query = strings.Replace(query, "\n", " ", -1)

	return strings.TrimSpace(query)
}
