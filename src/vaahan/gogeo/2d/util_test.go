package gogeo

import (
	"testing"
)

func TestRoundTo(t *testing.T) {
	testCases := []struct {
		input     float64
		precision int
		want      float64
	}{
		{
			input:     float64(5.323),
			precision: 2,
			want:      float64(5.32),
		},
		{
			input:     float64(5.423),
			precision: 2,
			want:      float64(5.42),
		},
		{
			input:     float64(5.325),
			precision: 2,
			want:      float64(5.33),
		},
		{
			input:     float64(5.425),
			precision: 2,
			want:      float64(5.43),
		},
		{
			input:     float64(5.327),
			precision: 2,
			want:      float64(5.33),
		},
		{
			input:     float64(5.427),
			precision: 2,
			want:      float64(5.43),
		},
		{
			input:     float64(-5.323),
			precision: 2,
			want:      float64(-5.32),
		},
		{
			input:     float64(-5.423),
			precision: 2,
			want:      float64(-5.42),
		},
		{
			input:     float64(-5.325),
			precision: 2,
			want:      float64(-5.33),
		},
		{
			input:     float64(-5.425),
			precision: 2,
			want:      float64(-5.43),
		},
		{
			input:     float64(-5.327),
			precision: 2,
			want:      float64(-5.33),
		},
		{
			input:     float64(-5.427),
			precision: 2,
			want:      float64(-5.43),
		},
	}

	for _, test := range testCases {
		result := RoundTo(test.input, test.precision)
		if result != test.want {
			t.Errorf("RoundTo(%v, %v) want: %v got: %v", test.input, test.precision, test.want, result)
		}
	}
}
