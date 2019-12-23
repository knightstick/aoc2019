package main

import (
	"testing"

	"github.com/knightstick/aoc2019/passwords"
)

func TestValidPassword(t *testing.T) {
	t.Run("all ones", func(t *testing.T) {
		assertValidPassword(t, 111111)
	})

	t.Run("descending digits", func(t *testing.T) {
		refuteValidPassword(t, 223450)
	})

	t.Run("descending at the start", func(t *testing.T) {
		refuteValidPassword(t, 911234)
	})

	t.Run("no double", func(t *testing.T) {
		refuteValidPassword(t, 123456)
	})
}

func assertValidPassword(t *testing.T, password int) {
	if !passwords.ValidPassword(password) {
		t.Errorf("expected %d to be valid", password)
	}
}

func refuteValidPassword(t *testing.T, password int) {
	if passwords.ValidPassword(password) {
		t.Errorf("expected %d not to be valid", password)
	}
}
