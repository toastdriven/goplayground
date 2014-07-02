package fizzbuzz

import "testing"

func TestCheckFizzBuzz(t *testing.T) {
	result := CheckFizzBuzz(1)

	if result != "1" {
		t.Error("Expected 1, got:", result)
	}

	result = CheckFizzBuzz(2)

	if result != "2" {
		t.Error("Expected 2, got:", result)
	}

	result = CheckFizzBuzz(3)

	if result != "Fizz" {
		t.Error("Expected Fizz, got:", result)
	}

	result = CheckFizzBuzz(4)

	if result != "4" {
		t.Error("Expected 4, got:", result)
	}

	result = CheckFizzBuzz(5)

	if result != "Buzz" {
		t.Error("Expected Buzz, got:", result)
	}

	result = CheckFizzBuzz(9)

	if result != "Fizz" {
		t.Error("Expected Fizz, got:", result)
	}

	result = CheckFizzBuzz(10)

	if result != "Buzz" {
		t.Error("Expected Buzz, got:", result)
	}

	result = CheckFizzBuzz(14)

	if result != "14" {
		t.Error("Expected 14, got:", result)
	}

	result = CheckFizzBuzz(15)

	if result != "FizzBuzz" {
		t.Error("Expected FizzBuzz, got:", result)
	}
}