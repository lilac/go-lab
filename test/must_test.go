package test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/mcesar/must"
)

func TestUseFile(t *testing.T) {
	assert.EqualError(t, useFile(), "open file: no such file or directory")
}

func useFile() (err error) {
	defer must.Handle(&err)
	f := must.Do(os.Open("file"))
	defer must.Do0(f.Close())
	return nil
}
