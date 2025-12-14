package formatter

import "go/format"

func FormatSource(src []byte) (formatted []byte, needsFormat bool, err error) {
	formatted, err = format.Source(src)
	if err != nil {
		return nil, false, err
	}

	needsFormat = string(formatted) != string(src)
	return formatted, needsFormat, nil
}
