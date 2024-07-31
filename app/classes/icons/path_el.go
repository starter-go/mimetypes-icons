package icons

import "strings"

func path2elements(path string) []string {
	const (
		s1 = "\\"
		s2 = "/"
	)
	path = strings.ReplaceAll(path, s1, s2)
	return strings.Split(path, s2)
}

func elements2path(pathEl []string) string {
	b := &strings.Builder{}
	for _, el := range pathEl {
		b.WriteRune('/')
		b.WriteString(el)
	}
	return b.String()
}

func normalizePathEl(pathEl []string) []string {
	dst := make([]string, 0)
	for _, el := range pathEl {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		} else if el == "." {
			continue
		} else if el == ".." {
			size := len(dst)
			if size > 0 {
				dst = dst[0 : size-1]
			} else {
				return []string{"[bad_path]"}
			}
		} else {
			dst = append(dst, el)
		}
	}
	return dst
}
