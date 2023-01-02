package nth_test

import (
	"fmt"
	"testing"

	"github.com/joeyak/nth"
)

func TestNth(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1st"},
		{input: 2, expected: "2nd"},
		{input: 3, expected: "3rd"},
		{input: 4, expected: "4th"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			actual := nth.Nth(tc.input)
			if tc.expected != actual {
				t.Errorf("expected %s got %s", tc.expected, actual)
			}
		})
	}
}

func TestGetNth(t *testing.T) {
	data := []int{1, 2, 3, 4}
	testCases := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1st"},
		{input: 2, expected: "2nd"},
		{input: 3, expected: "3rd"},
		{input: 4, expected: "4th"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			actual := nth.GetNth(tc.input, data)
			if tc.expected != actual {
				t.Errorf("expected %s got %s", tc.expected, actual)
			}
		})
	}
}

func TestNth0(t *testing.T) {
	defer func() {
		if r := recover(); r != nth.ErrZeroth {
			t.Errorf("expected %#v got %#v", nth.ErrZeroth, r)
		}
	}()
	nth.Nth(0)
}

func TestNthNegative(t *testing.T) {
	defer func() {
		if r := recover(); r != nth.ErrNegativeth {
			t.Errorf("expected %#v got %#v", nth.ErrNegativeth, r)
		}
	}()
	nth.Nth(-1)
}

func TestGetNthPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nth.ErrNoneth {
			t.Errorf("expected %#v got %#v", nth.ErrNoneth, r)
		}
	}()
	nth.GetNth(0, []int{1, 2, 3})
}
