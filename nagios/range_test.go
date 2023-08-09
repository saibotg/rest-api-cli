package nagios_test

import (
	"testing"

	"github.com/saibot/rest-api-cli/nagios"
)

func TestCheckValue(t *testing.T) {
	testCases := []struct {
		description string
		rule        string
		x           int
		want        bool
	}{
		{"-1 is outside range '10'", "10", -1, true},
		{"0 is not outside range '10'", "10", 0, false},
		{"5 is not outside range '10'", "10", 5, false},
		{"10 is not outside range '10'", "10", 10, false},
		{"11 is outside range '10'", "10", 11, true},
		{"2 is outside range '3:'", "3:", 2, true},
		{"3 is not outside range '3:'", "3:", 3, false},
		{"20 is not outside range '3:'", "3:", 20, false},
		{"-5 is not outside range '-10:'", "-10:", -5, false},
		{"23 is not outside range '10:30'", "10:30", 23, false},
		{"9 is outside range '10:30'", "10:30", 9, true},
		{"-2 is outside range '5:12", "5:12", -2, true},
		{"0 is outside range '-21:-5", "-21:-5", 0, true},
		{"-13 is not outside range '-21:-5", "-21:-5", -13, false},
		{"5 is outside range '~:0'", "~:0", 5, true},
		{"-10 is not outside range '~:10'", "~:10", -10, false},
		{"15 is inside range '@10:20'", "@10:20", 15, true},
		{"10 is inside range '@10:20'", "@10:20", 10, true},
		{"20 is inside range '@10:20'", "@10:20", 20, true},
		{"21 is not inside range '@10:20'", "@10:20", 21, false},
		{"9 is not inside range '@10:20'", "@10:20", 9, false},
		{"9 is not inside range '@10:'", "@10:", 9, false},
		{"11 is inside range '@10:'", "@10:", 11, true},
		{"-2 is inside range '@~:-1'", "@~:-1", -2, true},
		{"0 is not inside range '@~:-1'", "@~:-1", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := nagios.CheckValue(tc.rule, tc.x)
			if err != nil {
				t.Errorf("Error: %s\n", err)
			}
			if got != tc.want {
				t.Errorf("check value %d against rule '%s' should be %t, but was %t", tc.x, tc.rule, tc.want, got)
			}
		})
	}
}
