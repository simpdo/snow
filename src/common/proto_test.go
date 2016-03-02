package common

import (
	"testing"
)

func Test_Proto(t *testing.T) {
	head := Header{}
	head.Magic = MagicNum
	head.Cmd = 0x010001
	head.Version = 1
	head.BodyLen = 46

	data := EncodeHead(head)

	t.Log(data)

	resp := Header{}
	err := DecodeHead(data, &resp)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Magic != MagicNum {
		t.Error("magic decode failed")
		return
	}

	if resp.Cmd != head.Cmd {
		t.Error("cmd decode failed")
		return
	}

	if resp.Version != head.Version {
		t.Error("Version decode failed")
		return
	}

	if resp.BodyLen != head.BodyLen {
		t.Error("BodyLen decode failed")
		return
	}
}
