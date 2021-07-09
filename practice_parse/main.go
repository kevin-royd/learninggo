package main

import "strings"

func main() {
}

func Parse(rawurl string) (*URL, error) {
	// Cut off #frag
	u, frag := split(rawurl, '#', true)
	url, err := parse(u, false)
	if err != nil {
		return nil, &Error{"parse", u, err}
	}
	if frag == "" {
		return url, nil
	}
	if err = url.setFragment(frag); err != nil {
		return nil, &Error{"parse", rawurl, err}
	}
	return url, nil
}

func split(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}
