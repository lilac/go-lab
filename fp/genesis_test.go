package fp

import (
	"fmt"
	"github.com/life4/genesis/lambdas"
	"github.com/life4/genesis/slices"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenesis(t *testing.T) {
	b := lambdas.Must(slices.Min([]int{42, 7, 13})) == 7
	assert.True(t, b)

	items := []int{4, 8, 15}
	ints := slices.Map(items, func(el int) int { return el * 2 })
	assert.True(t, slices.Equal(ints, []int{8, 16, 30}))

	urls := []string{
		"https://go.dev/",
		"https://golang.org/",
		"https://google.com/",
	}
	codes := slices.MapAsync(
		urls, 0,
		func(url string) int {
			return lambdas.Must(http.Get(url)).StatusCode
		},
	)
	fmt.Println(codes)
	assert.True(t, slices.All(codes, func(el int) bool {
		return el == 200
	}))
}
