package tempconv

import "testing"

func TestCToF(t *testing.T) {
	tcs := []struct {
		temp    Celsius
		expects string
	}{
		{Celsius(0), "32°F"},
		{Celsius(32), "89.6°F"},
		{BoilingC, "212°F"},
	}

	for _, tc := range tcs {
		ret := CToF(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed CToF. Celsius: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}

func TestFToC(t *testing.T) {
	tcs := []struct {
		temp    Fahrenheit
		expects string
	}{
		{Fahrenheit(32), "0°C"},
		{Fahrenheit(212), "100°C"},
		{Fahrenheit(89.6), "32°C"},
	}

	for _, tc := range tcs {
		ret := FToC(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed FToC. Fahrenheit: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}

func TestCToK(t *testing.T) {
	tcs := []struct {
		temp    Celsius
		expects string
	}{
		{Celsius(0), "273.15°K"},
		{Celsius(100), "373.15°K"},
		{Celsius(32), "305.15°K"},
	}

	for _, tc := range tcs {
		ret := CToK(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed CToK. Celsius: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}

func TestKToC(t *testing.T) {
	tcs := []struct {
		temp    Kelvin
		expects string
	}{
		{Kelvin(273.15), "0°C"},
		{Kelvin(373.15), "100°C"},
		{Kelvin(305.15), "32°C"},
	}

	for _, tc := range tcs {
		ret := KToC(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed KToC. Kelvin: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}

func TestFToK(t *testing.T) {
	tcs := []struct {
		temp    Fahrenheit
		expects string
	}{
		{Fahrenheit(32), "273.15°K"},
		{Fahrenheit(212), "373.15°K"},
		{Fahrenheit(89.6), "305.15°K"},
	}

	for _, tc := range tcs {
		ret := FToK(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed FToK. Fahrenheit: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}

func TestKToF(t *testing.T) {
	tcs := []struct {
		temp    Kelvin
		expects string
	}{
		{Kelvin(273.15), "32°F"},
		{Kelvin(373.15), "212°F"},
		{Kelvin(305.15), "89.6°F"},
	}

	for _, tc := range tcs {
		ret := KToF(tc.temp)
		if ret.String() != tc.expects {
			t.Errorf("Failed KToF. Kelvin: %g, expects: %s, got %s", tc.temp, tc.expects, ret)
		}
	}
}
