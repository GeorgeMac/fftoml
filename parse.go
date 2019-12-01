package fftoml

import (
	"fmt"
	"io"
	"strings"

	"github.com/pelletier/go-toml"
)

// ParseConfigFile parses the provided reader as TOML
// and sets each found key on the provided set function
//
// It supports TOML tables by concatenating nested table keys
// together with a single "-" character
//
// For example, given the following document:
//
//  [config]
//  key = value
//  [config.property]
//  key = other-value
//
// The following flags can be set:
// flag.String("config-key", "", "some config flag")
// flag.String("config-property-key", "", "some other config flag")
func ParseConfigFile(r io.Reader, set func(name, value string) error) error {
	tree, err := toml.LoadReader(r)
	if err != nil {
		return err
	}

	return parseTree(tree, "", set)
}

func parseTree(tree *toml.Tree, parent string, set func(name, value string) error) error {
	for _, key := range tree.Keys() {
		name := key
		if parent != "" {
			name = parent + "-" + key
		}

		switch t := tree.Get(key).(type) {
		case *toml.Tree:
			if err := parseTree(t, name, set); err != nil {
				return err
			}
		case []interface{}:
			vs := []string{}
			for _, v := range t {
				vs = append(vs, fmt.Sprint(v))
			}

			if err := set(name, strings.Join(vs, ",")); err != nil {
				return err
			}
		case interface{}:
			if err := set(name, fmt.Sprint(t)); err != nil {
				return err
			}
		}
	}

	return nil
}
