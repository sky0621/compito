package util

// Separate ...
func Separate(chars []string, idx int) (string, []string) {
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
