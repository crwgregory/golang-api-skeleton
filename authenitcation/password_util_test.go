package authentication

import (
	"testing"
)

type cryptTest struct {
	password string
	expected string
}

var (
	test_key = []byte{0xe3, 0x1f, 0x4d, 0xa2, 0xaa, 0x11, 0x00, 0x12, 0x01}
	tests    = []cryptTest{
		{
			"123456789",
			"LeRNRMaUS2CiqRLxhpOPczbDiBw1gBY5piCSeHGOYxe3hVpY6nKDDBaRsJqt8r3Q",
		},
		{
			"abcdefghijklmnopqrstuv",
			"D5zOkk/IUy0hCKh7X3MAbI031yiOldbZq3PRc4cJ2UsGVXXnLrALoPeOVJQdKWG9",
		},
		{
			"ASLDFKJlakjsdf*%@*",
			"o8IqT1+0lxy0brRpu1AFK0b5mwTdHVeCSgA2fYPGIBZFP0bDzRLjJ5Ka7giVQiZH",
		},
		{
			"nvakdfh*^@))!@#$%^&*()_",
			"g7d9nnrV9dXCj1xYGe7hNFrhsVaPOcVIwpVUntTxvvTOKd4bRR+TNvPG5oWHtyTA",
		},
		{
			"12345678",
			"m+G8HAHLYKqJ6dDAEEG9pXEpAX1NQ2CTUCHKF4luPSN9rXCADJxVyi/tMIzUN+x8",
		},
		{
			"1234567'\"8",
			"EIckzD57ARjB4naNL4Qb9jQc//lcuWaH4i9tswat/QMhMmK9uqmYKHAQbrzdjWrP",
		},
	}
)

func TestPasswordHasher_HashPassword(t *testing.T) {
	hasher := new(PasswordHasher)
	hasher.Key = test_key

	for _, test := range tests {
		result, err := hasher.HashPassword(test.password)

		if err != nil {
			t.Log(test)
			t.Error(err)
		}

		if len(result) != len(test.expected) {
			t.Error("result is not same length as expected")
		}

		if result != test.expected {
			t.Error("result is not equal to expected", result, test.expected)
		}
	}
}
