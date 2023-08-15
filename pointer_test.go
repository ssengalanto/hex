package hex_test

import (
	"testing"
	"time"

	"github.com/ssengalanto/hex"
	"github.com/stretchr/testify/assert"
)

func TestDeref_NilPointer(t *testing.T) {
	var nilPtr *int
	derefNilValue := hex.Deref(nilPtr)
	assert.Equal(t, 0, derefNilValue)
}

func TestDeref_Int(t *testing.T) {
	input := (*int)(nil)
	expected := 0
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}

func TestDeref_String(t *testing.T) {
	input := func() *string { val := "hello"; return &val }()
	expected := "hello"
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}

func TestDeref_Bool(t *testing.T) {
	input := (*bool)(nil)
	actual := hex.Deref(input)
	assert.Equal(t, false, actual)
}

func TestDeref_Uint(t *testing.T) {
	input := (*uint)(nil)
	expected := uint(0)
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}

func TestDeref_Float64(t *testing.T) {
	input := (*float64)(nil)
	expected := float64(0)
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}

func TestDeref_IntSlice(t *testing.T) {
	input := (*[]int)(nil)
	expected := []int(nil)
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}

func TestDeref_Time(t *testing.T) {
	input := (*time.Time)(nil)
	expected := time.Time{}
	actual := hex.Deref(input)
	assert.Equal(t, expected, actual)
}
