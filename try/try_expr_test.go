package main

import (
	"encoding/json"
	"os"
	"testing"
)
import trye "github.com/dsnet/try"

func Test(t *testing.T) {
	defer trye.F(t.Fatal)
	b := trye.E1(os.ReadFile("/tmp/a.log"))
	var v any
	trye.E(json.Unmarshal(b, &v))
}
