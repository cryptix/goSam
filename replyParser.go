package goSam

import (
	"fmt"
	"strings"
)

type Reply struct {
	Topic string
	Type  string

	Pairs map[string]string
}

func parseReply(line string) (r *Reply, err error) {
	parts := strings.Split(line, " ")
	if len(parts) < 3 {
		return nil, fmt.Errorf("Malformed Reply.\n%s\n", line)
	}

	r = &Reply{
		Topic: parts[0],
		Type:  parts[1],
		Pairs: make(map[string]string, len(parts)-2),
	}

	for _, v := range parts[2:] {
		kvPair := strings.Split(v, "=")
		if len(kvPair) != 2 {
			return nil, fmt.Errorf("Malformed key-value-pair.\n%s\n", kvPair)
		}
		r.Pairs[kvPair[0]] = kvPair[1]
	}

	return
}
