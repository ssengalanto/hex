package hex_test

import (
	"testing"

	"github.com/ssengalanto/hex"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Parallel()

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
			t.Parallel()
			v, ok := hex.Find(tc.s, tc.fn)
			assert.Equal(t, tc.want, v)
			assert.Equal(t, tc.found, ok)
		})
	}
}

func TestFindIndex(t *testing.T) {
	t.Parallel()

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
			t.Parallel()
			idx := hex.FindIndex(tc.s, tc.fn)
			assert.Equal(t, tc.want, idx)
		})
	}
}
