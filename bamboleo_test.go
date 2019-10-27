package bamboleo

import (
	"fmt"
	"testing"
)

func TestBamboleo(t *testing.T) {
	for _, test := range stringTestCases {
		bamboleo := Parse(test.input)
		fmt.Print(bamboleo)
		// if actual != test.expected {
		// 	t.Errorf("Bamboleo test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		// }
	}
}

func BenchmarkBamboleo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Parse(test.input)
		}
	}
}

func TestBamboleoJSON(t *testing.T) {
	for _, test := range stringTestCases {
		bamboleo := ParseToJSON(test.input)
		fmt.Print(bamboleo)
		// if actual != test.expected {
		// 	t.Errorf("Bamboleo test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		// }
	}
}

func BenchmarkBamboleoJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			ParseToJSON(test.input)
		}
	}
}
