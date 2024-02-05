package tests

import (
	"testing"

	"github.com/AlbertRossJoh/itualgs_go/utilities"
)

func TestOptionalNoValue(t *testing.T) {
	option := new(utilities.Optional[int])
	var v *int
	option.Some(v)
	if v != nil {
		t.Error("Expected nil pointer found:", *v)
	}
}

func TestOptionalHasValue(t *testing.T) {
	option := utilities.NewOptional(2)
	var v int = 0
	option.Some(&v)
	if v == 0 {
		t.Error("Expected value pointer found nil")
	}
}
