package testProperty

import (
	"fortunate_street/property"
	"testing"
)

// tests
func TestCalcCost(t *testing.T) {
	testProp := property.Property{Name: "test", BaseCost: 400, Investment: 100}
	if property.CalcCost(testProp) != 500 {
		t.Errorf("oops")
	}
}
