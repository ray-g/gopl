package main

import "testing"
import "bytes"

func TestTempUnit(t *testing.T) {
	tcs := []struct {
		number  float64
		expects string
	}{
		{1.0, "1.00°F = -17.22°C, 1.00°C = 33.80°F\n"},
		{0.0, "0.00°F = -17.78°C, 0.00°C = 32.00°F\n"},
		{-1.0, "-1.00°F = -18.33°C, -1.00°C = 30.20°F\n"},
	}

	for _, tc := range tcs {
		ret := tempUnit(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed tempUnit. Number: %g, expects: %s, got %s", tc.number, tc.expects, ret)
		}
	}
}

func TestLenUnit(t *testing.T) {
	tcs := []struct {
		number  float64
		expects string
	}{
		{1.0, "1.00 FT = 0.30 M, 1.00 M = 3.28 FT\n"},
		{0.0, "0.00 FT = 0.00 M, 0.00 M = 0.00 FT\n"},
		{-1.0, "-1.00 FT = -0.30 M, -1.00 M = -3.28 FT\n"},
	}

	for _, tc := range tcs {
		ret := lenUnit(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed lenUnit. Number: %g, expects: %s, got %s", tc.number, tc.expects, ret)
		}
	}
}

func TestWeightUnit(t *testing.T) {
	tcs := []struct {
		number  float64
		expects string
	}{
		{1.0, "1.00 P = 0.45 KG, 1.00 KG = 2.20 P\n"},
		{0.0, "0.00 P = 0.00 KG, 0.00 KG = 0.00 P\n"},
		{-1.0, "-1.00 P = -0.45 KG, -1.00 KG = -2.20 P\n"},
	}

	for _, tc := range tcs {
		ret := weightUnit(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed weightUnit. Number: %g, expects: %s, got %s", tc.number, tc.expects, ret)
		}
	}
}

func TestAllUnit(t *testing.T) {
	tcs := []struct {
		number  float64
		expects string
	}{
		{1.0, "1.00°F = -17.22°C, 1.00°C = 33.80°F\n1.00 FT = 0.30 M, 1.00 M = 3.28 FT\n1.00 P = 0.45 KG, 1.00 KG = 2.20 P\n"},
		{0.0, "0.00°F = -17.78°C, 0.00°C = 32.00°F\n0.00 FT = 0.00 M, 0.00 M = 0.00 FT\n0.00 P = 0.00 KG, 0.00 KG = 0.00 P\n"},
		{-1.0, "-1.00°F = -18.33°C, -1.00°C = 30.20°F\n-1.00 FT = -0.30 M, -1.00 M = -3.28 FT\n-1.00 P = -0.45 KG, -1.00 KG = -2.20 P\n"},
	}

	for _, tc := range tcs {
		ret := allUnits(tc.number)
		if ret != tc.expects {
			t.Errorf("Failed allUnits. Number: %g, expects: %s, got %s", tc.number, tc.expects, ret)
		}
	}
}

func TestPrintAll(t *testing.T) {
	tcs := []struct {
		nums    []string
		expects string
	}{
		{[]string{"1.0", "0.0", "-1.0"},
			"1.00°F = -17.22°C, 1.00°C = 33.80°F\n" +
				"1.00 FT = 0.30 M, 1.00 M = 3.28 FT\n" +
				"1.00 P = 0.45 KG, 1.00 KG = 2.20 P\n" +
				"----------\n" +
				"0.00°F = -17.78°C, 0.00°C = 32.00°F\n" +
				"0.00 FT = 0.00 M, 0.00 M = 0.00 FT\n" +
				"0.00 P = 0.00 KG, 0.00 KG = 0.00 P\n" +
				"----------\n" +
				"-1.00°F = -18.33°C, -1.00°C = 30.20°F\n" +
				"-1.00 FT = -0.30 M, -1.00 M = -3.28 FT\n" +
				"-1.00 P = -0.45 KG, -1.00 KG = -2.20 P\n" +
				"----------\n"},
	}

	for _, tc := range tcs {
		stdout = new(bytes.Buffer)
		printAll(tc.nums)
		if stdout.(*bytes.Buffer).String() != tc.expects {
			t.Errorf("Failed printAll. Numbers: %q, expects: %s, got %s", tc.nums, tc.expects, stdout)
		}
	}
}
