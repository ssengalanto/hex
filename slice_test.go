package hex_test

import (
	"testing"

	"github.com/ssengalanto/hex"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		fn    func(n int) bool
		found bool
		want  int
	}{
		{
			name: "element found",
			s:    []int{1, 2, 3},
			fn: func(n int) bool {
				return n > 2
			},
			found: true,
			want:  3,
		},
		{
			name: "element not found",
			s:    []int{3, 2, 1},
			fn: func(n int) bool {
				return n < 1
			},
			found: false,
			want:  0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v, ok := hex.Find(tc.s, tc.fn)
			assert.Equal(t, tc.want, v)
			assert.Equal(t, tc.found, ok)
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		fn   func(n int) bool
		want int
	}{
		{
			name: "element found",
			s:    []int{1, 2, 3},
			fn: func(n int) bool {
				return n == 2
			},
			want: 1,
		},
		{
			name: "element not found",
			s:    []int{1, 2, 3},
			fn: func(n int) bool {
				return n == 0
			},
			want: -1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			idx := hex.FindIndex(tc.s, tc.fn)
			assert.Equal(t, tc.want, idx)
		})
	}
}

func TestMap(t *testing.T) {
	t.Run("returns the same slice type", func(t *testing.T) {
		s := hex.Map([]int{1, 2, 3}, func(el int, i int) int {
			return el * el
		})
		assert.Equal(t, []int{1, 4, 9}, s)
	})

	t.Run("returns different slice type", func(t *testing.T) {
		s := hex.Map([]int{104, 101, 108, 108, 111}, func(el int, i int) string {
			return string(rune(el))
		})
		assert.Equal(t, []string{"h", "e", "l", "l", "o"}, s)
	})
}

func TestFilter(t *testing.T) {
	t.Run("returns filtered slice", func(t *testing.T) {
		s := hex.Filter([]int{1, 2, 3, 4}, func(el int, i int) bool {
			return el > 2
		})
		assert.Equal(t, []int{3, 4}, s)
	})

	t.Run("returns empty slice", func(t *testing.T) {
		s := hex.Filter([]int{1, 2, 3, 4}, func(el int, i int) bool {
			return el > 5
		})
		assert.Equal(t, []int{}, s)
	})
}

func TestChunk(t *testing.T) {
	t.Run("returns chunked slice", func(t *testing.T) {
		s := hex.Chunk([]int{1, 2, 3, 4, 5}, 2)
		assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, s)
	})

	t.Run("panics when size less than 1", func(t *testing.T) {
		assert.Panics(t, func() {
			hex.Chunk([]int{1, 2, 3, 4, 5}, 0)
		})
	})
}

func TestPrepend(t *testing.T) {
	t.Run("prepends the element in the slice", func(t *testing.T) {
		s := hex.Prepend([]int{1, 2, 3}, 0)
		assert.Equal(t, []int{0, 1, 2, 3}, s)
	})
}
