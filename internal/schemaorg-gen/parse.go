package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type triple struct {
	Subj, Pred, Obj string
}

func parseNT(r io.Reader) ([]triple, error) {
	var out []triple
	sc := bufio.NewScanner(r)
	const max = 1 << 30
	sc.Buffer(make([]byte, 0, 64*1024), max)
	lineNum := 0
	for sc.Scan() {
		lineNum++
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		t, err := parseNTLine(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum, err)
		}
		if t != nil {
			out = append(out, *t)
		}
	}
	return out, sc.Err()
}

func parseNTLine(line string) (*triple, error) {
	if !strings.HasSuffix(line, " .") {
		return nil, fmt.Errorf("expected line to end with ' .'")
	}
	line = strings.TrimSuffix(line, " .")
	line = strings.TrimSpace(line)

	subj, rest, err := nextToken(line)
	if err != nil {
		return nil, fmt.Errorf("subject: %w", err)
	}
	pred, rest, err := nextToken(strings.TrimSpace(rest))
	if err != nil {
		return nil, fmt.Errorf("predicate: %w", err)
	}
	obj, rest, err := nextToken(strings.TrimSpace(rest))
	if err != nil {
		return nil, fmt.Errorf("object: %w", err)
	}
	if strings.TrimSpace(rest) != "" {
		return nil, fmt.Errorf("trailing %q", strings.TrimSpace(rest))
	}
	return &triple{Subj: subj, Pred: pred, Obj: obj}, nil
}

// nextToken returns an IRI <...> or literal "..." (with escapes) and remaining string.
func nextToken(s string) (tok, rest string, err error) {
	s = strings.TrimLeft(s, " \t")
	if s == "" {
		return "", "", fmt.Errorf("empty token")
	}
	if s[0] == '<' {
		i := strings.IndexByte(s[1:], '>')
		if i < 0 {
			return "", "", fmt.Errorf("unclosed IRI")
		}
		return s[1 : i+1], s[i+2:], nil
	}
	if s[0] == '"' {
		i := 1
		var b strings.Builder
		for i < len(s) {
			c := s[i]
			if c == '\\' && i+1 < len(s) {
				i++
				switch s[i] {
				case 'n':
					b.WriteByte('\n')
				case 't':
					b.WriteByte('\t')
				case 'r':
					b.WriteByte('\r')
				case '"', '\\':
					b.WriteByte(s[i])
				default:
					b.WriteByte(s[i])
				}
				i++
				continue
			}
			if c == '"' {
				rest = s[i+1:]
				if strings.HasPrefix(rest, "@") {
					if sp := strings.IndexByte(rest, ' '); sp >= 0 {
						rest = rest[sp:]
					} else {
						rest = ""
					}
				} else if strings.HasPrefix(strings.TrimLeft(rest, " \t"), "^^") {
					rest = strings.TrimLeft(rest, " \t")
					rest = rest[2:]
					_, rest, err = nextToken(strings.TrimLeft(rest, " \t"))
					if err != nil {
						return "", "", err
					}
				}
				return b.String(), strings.TrimLeft(rest, " \t"), nil
			}
			b.WriteByte(c)
			i++
		}
		return "", "", fmt.Errorf("unclosed string literal")
	}
	return "", "", fmt.Errorf("unexpected %q", s[:min(12, len(s))])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
