package word

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Errorf(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Errorf(`IsPalindrome("kayak") = false`)
	}
}

func TestNonePalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Errorf(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Errorf(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
