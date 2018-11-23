package util

// Separate ...
func Separate(chars []string, idx int) (string, []string) {
	if idx < 0 {
		return "", []string{}
	}
	if chars == nil {
		return "", []string{}
	}
	if len(chars) == 0 {
		return "", []string{}
	}
	if len(chars) <= idx {
		return "", []string{}
	}
	var target string
	var remaining []string
	for i, c := range chars {
		if i == idx {
			target = c
		} else {
			remaining = append(remaining, c)
		}
	}
	return target, remaining
}
