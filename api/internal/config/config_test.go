package config

import (
	"testing"

	"io.github.tuanictu97/api/pkg/util/toml"
)

func TestConfig(t *testing.T) {
	MustLoad("../../configs/config.toml")

	s, err := toml.MarshalToString(C)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n", s, "\n")
}
