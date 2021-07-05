package main

import (
	"testing"
)

func TestNetting(t *testing.T) {
	s := &Spider{
		Url:      "https://bbs.hupu.com/",
		Name:     "bbs_hupu_com",
		Timeout:  30,
		Children: nil,
	}
	Netting(s)
}
