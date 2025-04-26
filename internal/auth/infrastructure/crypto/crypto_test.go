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
