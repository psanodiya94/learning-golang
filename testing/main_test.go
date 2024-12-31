package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expected-data", 100.0, 25.0, 4.0, false},
	{"expected-frac", -1.0, -777, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		result, err := devide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one")
			}
		}

		if result != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, result)
		}
	}
}
