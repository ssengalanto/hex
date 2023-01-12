package hex_test

import (
	"testing"

	"github.com/ssengalanto/hex"
	"github.com/stretchr/testify/assert"
)

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
