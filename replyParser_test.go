package goSam

import (
	"testing"
)

var validCases = []struct {
	Input    string
	Expected Reply
}{
	// hello handshake reply
	{
		"HELLO REPLY RESULT=OK VERSION=3.0\n",
		Reply{
			Topic: "HELLO",
			Type:  "REPLY",
			Pairs: map[string]string{
				"RESULT":  "OK",
				"VERSION": "3.0",
			},
		},
	},
	// result of a naming lookup
	{
		"NAMING REPLY RESULT=OK NAME=zzz.i2p VALUE=SomeValueForTesting\n",
		Reply{
			Topic: "NAMING",
			Type:  "REPLY",
			Pairs: map[string]string{
				"RESULT": "OK",
				"NAME":   "zzz.i2p",
				"VALUE":  "SomeValueForTesting",
			},
		},
	},
	// result of a b32.i2p naming lookup
	{
		"NAMING REPLY RESULT=OK NAME=gkso46tc47hdua2kva5zahj3unmyh6ia7bv5oc2ybn2hmeowpz7a.b32.i2p VALUE=mlHQraXLpcE7A4MVeVniRHM~2yoaW1fOYVKj3ZiNTe4UPIAlIReYUMHSnZnloFWX7bh2OoEg08JBGoSQPtGkCZjqSBmfeDdMqtwbZ~~sok-jNo4e5rWnfCOHYYPVcuE8jB~7M5ioJzk8QZRqh3AjCQsKBUZgTzUfGlP12j3GtAf5C9iAGxTTB1sGE96752EKP0dzyGOs4NAujwkgm6NzVoqlkXD~fognUrQOeG~~OqChsAeqIRqj40FsJmsJREmZ4GhjFAqzLUQ4hMpKSbqMI~wtfjeNs-WKtM7pCD09uwSmYwW84wu-WxLGZiIt2GKmbPv~JrqYFNv9EM0SzUonAF8pw9GAhUn8-26kkgCXTs05Kut7NuQHghu3jHfS-frlPmAt-Uu5T4ZcLiHiFrnG2lYTtjxBFXh7W72IBncHSixhVhd4lYM7REKFj7G~5ttW9EBeL1unbNYWiQOEQjtGlmwxYt~~2EV16w339aQQ~S~69-tS6vFA1n2DgkMdg06pBQAEAAEAAA==\n",
		Reply{
			Topic: "NAMING",
			Type:  "REPLY",
			Pairs: map[string]string{
				"RESULT": "OK",
				"NAME":   "gkso46tc47hdua2kva5zahj3unmyh6ia7bv5oc2ybn2hmeowpz7a.b32.i2p",
				"VALUE":  "mlHQraXLpcE7A4MVeVniRHM~2yoaW1fOYVKj3ZiNTe4UPIAlIReYUMHSnZnloFWX7bh2OoEg08JBGoSQPtGkCZjqSBmfeDdMqtwbZ~~sok-jNo4e5rWnfCOHYYPVcuE8jB~7M5ioJzk8QZRqh3AjCQsKBUZgTzUfGlP12j3GtAf5C9iAGxTTB1sGE96752EKP0dzyGOs4NAujwkgm6NzVoqlkXD~fognUrQOeG~~OqChsAeqIRqj40FsJmsJREmZ4GhjFAqzLUQ4hMpKSbqMI~wtfjeNs-WKtM7pCD09uwSmYwW84wu-WxLGZiIt2GKmbPv~JrqYFNv9EM0SzUonAF8pw9GAhUn8-26kkgCXTs05Kut7NuQHghu3jHfS-frlPmAt-Uu5T4ZcLiHiFrnG2lYTtjxBFXh7W72IBncHSixhVhd4lYM7REKFj7G~5ttW9EBeL1unbNYWiQOEQjtGlmwxYt~~2EV16w339aQQ~S~69-tS6vFA1n2DgkMdg06pBQAEAAEAAA==",
			},
		},
	},
	// session status reply
	{
		"SESSION STATUS RESULT=I2P_ERROR MESSAGE=TheMessageFromI2p\n",
		Reply{
			Topic: "SESSION",
			Type:  "STATUS",
			Pairs: map[string]string{
				"RESULT":  "I2P_ERROR",
				"MESSAGE": "TheMessageFromI2p",
			},
		},
	},
	{
		"STREAM STATUS RESULT=CANT_REACH_PEER\n",
		Reply{
			Topic: "STREAM",
			Type:  "STATUS",
			Pairs: map[string]string{
				"RESULT": "CANT_REACH_PEER",
			},
		},
	},
}

func TestParseReplyValidCases(t *testing.T) {
	for _, tcase := range validCases {
		parsed, err := parseReply(tcase.Input)
		if err != nil {
			t.Fatalf("parseReply should not throw an error!\n%s", err)
		}

		if parsed.Topic != tcase.Expected.Topic {
			t.Fatalf("Wrong Topic. Got %s expected %s", parsed.Topic, tcase.Expected.Topic)
		}

		if len(parsed.Pairs) != len(tcase.Expected.Pairs) {
			t.Fatalf("Wrong ammount of Pairs. Got %d expected 3", len(parsed.Pairs))
		}

		for expK, expV := range tcase.Expected.Pairs {
			if expV != parsed.Pairs[expK] {
				t.Fatalf("Wrong %s.\nGot '%s'\nExpected '%s'", expK, parsed.Pairs[expK], expV)
			}
		}
	}
}

func TestParseInvalidReply(t *testing.T) {
	str := "asd asd="

	r, err := parseReply(str)
	if err == nil {
		t.Fatalf("Should throw an error.r:%v\n", r)
	}
}
