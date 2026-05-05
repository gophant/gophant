package templates

import (
	"io/fs"
)

// ListEmbeddedTemplates returns a map of architecture -> variants available in embedded templates.
// For architectures that only provide files at the root, the variant name will be "default".
func ListEmbeddedTemplates() (map[string][]string, error) {
	out := make(map[string][]string)

	// helper to inspect an embed FS at base path
	check := func(fsys fs.FS, base string) ([]string, error) {
		entries, err := fs.ReadDir(fsys, base)
		if err != nil {
			return nil, err
		}
		var variants []string
		for _, e := range entries {
			if e.IsDir() {
				variants = append(variants, e.Name())
			}
		}
		// if no subdirs but entries exist, treat as single default variant
		if len(variants) == 0 && len(entries) > 0 {
			variants = []string{"default"}
		}
		return variants, nil
	}

	if v, err := check(defaultFS, "default"); err == nil {
		out["default"] = v
	}
	if v, err := check(mvcFS, "mvc"); err == nil {
		out["mvc"] = v
	}
	if v, err := check(dddFS, "ddd"); err == nil {
		out["ddd"] = v
	}

	return out, nil
}
