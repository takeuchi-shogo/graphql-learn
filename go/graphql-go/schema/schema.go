package schema

import (
	"bytes"
	"embed"
	"errors"
	"io/fs"
	"strings"
)

//go:embed graphql/schema.gql
var content embed.FS

func String() (string, error) {
	var buf bytes.Buffer

	fn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.New("walking dir")
		}

		if !strings.HasSuffix(path, ".gql") {
			return nil
		}

		b, err := content.ReadFile(path)
		if err != nil {
			return errors.New("reading file " + path)
		}

		b = append(b, []byte("\n")...)

		if _, err := buf.Write(b); err != nil {
			return errors.New("writing " + path + " bytes to buffer")
		}

		return nil
	}

	if err := fs.WalkDir(content, ".", fn); err != nil {
		return buf.String(), errors.New("walking content directory")
	}

	return buf.String(), nil
}
