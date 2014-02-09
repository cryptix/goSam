package goSam

import (
	"testing"
)

func TestParseInvalidReply(t *testing.T) {
	str := "asd asd="

	r, err := parseReply(str)
	if err == nil {
		t.Fatalf("Should throw an error.r:%v\n", r)
	}
}

func TestParseReplyNAMING(t *testing.T) {
	str := "NAMING REPLY RESULT=OK NAME=zzz.i2p VALUE=GKapJ8koUcBj~jmQzHsTYxDg2tpfWj0xjQTzd8BhfC9c3OS5fwPBNajgF-eOD6eCjFTqTlorlh7Hnd8kXj1qblUGXT-tDoR9~YV8dmXl51cJn9MVTRrEqRWSJVXbUUz9t5Po6Xa247Vr0sJn27R4KoKP8QVj1GuH6dB3b6wTPbOamC3dkO18vkQkfZWUdRMDXk0d8AdjB0E0864nOT~J9Fpnd2pQE5uoFT6P0DqtQR2jsFvf9ME61aqLvKPPWpkgdn4z6Zkm-NJOcDz2Nv8Si7hli94E9SghMYRsdjU-knObKvxiagn84FIwcOpepxuG~kFXdD5NfsH0v6Uri3usE3uSzpWS0EHmrlfoLr5uGGd9ZHwwCIcgfOATaPRMUEQxiK9q48PS0V3EXXO4-YLT0vIfk4xO~XqZpn8~PW1kFe2mQMHd7oO89yCk-3yizRG3UyFtI7-mO~eCI6-m1spYoigStgoupnC3G85gJkqEjMm49gUjbhfWKWI-6NwTj0ZnAAAA"

	reply, err := parseReply(str)
	if err != nil {
		t.Fatalf("parseReply should not throw an error!\n%s", err)
	}

	if reply.Topic != "NAMING" {
		t.Fatalf("Wrong Topic. Got %s expected NAMING", reply.Topic)
	}

	if reply.Type != "REPLY" {
		t.Fatalf("Wrong Type. Got %s expected REPLY", reply.Type)
	}

	if len(reply.Pairs) != 3 {
		t.Fatalf("Wrong ammount of Pairs. Got %d expected 3", len(reply.Pairs))
	}

	for k, v := range reply.Pairs {
		switch k {
		case "RESULT":
			if v != "OK" {
				t.Fatalf("Wrong Result. Got %s expected OK", v)
			}
		case "NAME":
			if v != "zzz.i2p" {
				t.Fatalf("Wrong Name. Got %s expected OK", v)
			}
		case "VALUE":
			expect := "GKapJ8koUcBj~jmQzHsTYxDg2tpfWj0xjQTzd8BhfC9c3OS5fwPBNajgF-eOD6eCjFTqTlorlh7Hnd8kXj1qblUGXT-tDoR9~YV8dmXl51cJn9MVTRrEqRWSJVXbUUz9t5Po6Xa247Vr0sJn27R4KoKP8QVj1GuH6dB3b6wTPbOamC3dkO18vkQkfZWUdRMDXk0d8AdjB0E0864nOT~J9Fpnd2pQE5uoFT6P0DqtQR2jsFvf9ME61aqLvKPPWpkgdn4z6Zkm-NJOcDz2Nv8Si7hli94E9SghMYRsdjU-knObKvxiagn84FIwcOpepxuG~kFXdD5NfsH0v6Uri3usE3uSzpWS0EHmrlfoLr5uGGd9ZHwwCIcgfOATaPRMUEQxiK9q48PS0V3EXXO4-YLT0vIfk4xO~XqZpn8~PW1kFe2mQMHd7oO89yCk-3yizRG3UyFtI7-mO~eCI6-m1spYoigStgoupnC3G85gJkqEjMm49gUjbhfWKWI-6NwTj0ZnAAAA"
			if v != expect {
				t.Fatalf("Wrong Value.\nGot:%s\nExpected:%s", v, expect)
			}
		default:
			t.Fatalf("Unknown kvPair %s=%s", k, v)
		}
	}
}

func TestParseReplyHELLO(t *testing.T) {
	str := "HELLO REPLY RESULT=OK VERSION=3.0"

	reply, err := parseReply(str)
	if err != nil {
		t.Fatalf("parseReply should not throw an error!\n%s", err)
	}

	if reply.Topic != "HELLO" {
		t.Fatalf("Wrong Topic. Got %s expected HELLO", reply.Topic)
	}

	if reply.Type != "REPLY" {
		t.Fatalf("Wrong Type. Got %s expected REPLY", reply.Type)
	}

	if len(reply.Pairs) != 2 {
		t.Fatalf("Wrong ammount of Pairs. Got %d expected 3", len(reply.Pairs))
	}

	for k, v := range reply.Pairs {
		switch k {
		case "RESULT":
			if v != "OK" {
				t.Fatalf("Wrong Result. Got %s expected OK", v)
			}
		case "VERSION":
			if v != "3.0" {
				t.Fatalf("Wrong Name. Got %s expected OK", v)
			}
		default:
			t.Fatalf("Unknown kvPair %s=%s", k, v)
		}
	}
}
