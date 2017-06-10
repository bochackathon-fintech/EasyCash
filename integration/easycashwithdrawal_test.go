package integration

import (
	"testing"
)

func TestCount(t *testing.T) {
	if GetBalance2() != "5710bba5d42604e4072d1e92" {
		t.Error("Expected 5710bba5d42604e4072d1e92")
	}
}
