package main

import "testing"

func TestMultiple(t *testing.T) {
	// preparing data for tests, extra logic, db connection, etc.
	t.Run("simple", func(t *testing.T) {
		var a, b, result = 2, 4, 8

		res := multiple(a, b)

		if res != result {
			t.Errorf("%d != %d", res, result)
		}
	})
	t.Run("medium", func(t *testing.T) {
		var a, b, result = 232, 4, 928

		res := multiple(a, b)

		if res != result {
			t.Errorf("%d != %d", res, result)
		}
	})
	t.Run("negative", func(t *testing.T) {
		var a, b, result = -2, 4, -8

		res := multiple(a, b)

		if res != result {
			t.Errorf("%d != %d", res, result)
		}
	})
	// removing data after tests
}
