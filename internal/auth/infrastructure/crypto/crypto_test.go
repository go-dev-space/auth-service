package crypto

import "testing"

func Test_RandomString(t *testing.T) {

	tests := []struct {
		name           string
		length         int
		expectedResult int
	}{
		{"Valid string", 7, 7},
	}

	for _, test := range tests {
		newString := New().GenerateRandomString(test.length)

		t.Run(test.name, func(t *testing.T) {
			if test.expectedResult != len(newString) {
				t.Errorf("[%s] expected length %d, but string length is %d", test.name, test.expectedResult, len(newString))
			}
		})
	}

}

func Test_RandomInt(t *testing.T) {

	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Valid number", 1, 10},
	}

	for _, test := range tests {
		rnd := New().GenerateRandomInt(test.min, test.max)

		t.Run(test.name, func(t *testing.T) {
			if rnd > test.max || rnd < test.min {
				t.Errorf("[%s] random number is out of range! Expected rang %d-%d, but go %d", test.name, test.min, test.max, rnd)
			}
		})
	}

}
