package htmlmetaparse

import "strings"

func Parse(s string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	for {
		idx1 := strings.Index(s, "<meta ")
		if idx1 == -1 {
			break
		}
		idx2 := strings.Index(s[idx1+5:], ">")
		if idx2 == -1 {
			break
		}

		ss := s[idx1 : idx1+5+idx2+1]
		println(ss)
		p, _ := parseTag(ss)
		res[ss] = p
		s = s[idx1+5+idx2+1:]
	}
	return res, nil
}

func parseTag(s string) (map[string]string, error) {
	s = strings.TrimPrefix(s, "<meta")
	s = strings.TrimSuffix(s, ">")

	isKey, isValue := false, false
	var key, val string

	res := map[string]string{}
	for _, ch := range s {
		switch {
		default:
			if isKey {
				key += string(ch)
			} else if isValue {
				val += string(ch)
			} else {
				isKey = true
				key += string(ch)
			}
		case ch == '=':
			if isKey {
				isKey = false
			} else if isValue {
				val += string(ch)
			}
		case ch == '"':
			if !isKey && !isValue {
				isValue = true
			} else if isValue {
				isValue = false
			}
		case ch == ' ':
			if isValue {
				val += string(ch)
				continue
			}
			if key == "" {
				continue
			}
			res[key] = val
			key, val = "", ""
			isKey, isValue = false, false
		}
	}
	res[key] = val
	return res, nil
}
