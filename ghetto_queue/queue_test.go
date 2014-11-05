package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := Queue{}

	if queue.length() != 0 {
		t.Errorf("Expected zero length queue, was:", queue.length())
	}

	queue.put("Hello")
	queue.put("World")
	queue.put("foo")

	if queue.length() != 3 {
		t.Errorf("Expected 3 length queue, was:", queue.length())
	}

	if queue.empty() != false {
		t.Errorf("Expected non-empty queue, was:", queue.empty())
	}

	first := queue.get()

	if first != "Hello" {
		t.Errorf("Expected \"Hello\", was:", first)
	}

	if queue.length() != 2 {
		t.Errorf("Expected 2 length queue, was:", queue.length())
	}

	second := queue.get()

	if second != "World" {
		t.Errorf("Expected \"World\", was:", second)
	}

	if queue.empty() != false {
		t.Errorf("Expected non-empty queue, was:", queue.empty())
	}

	third := queue.get()

	if third != "foo" {
		t.Errorf("Expected \"foo\", was:", third)
	}

	if queue.length() != 0 {
		t.Errorf("Expected 0 length queue, was:", queue.length())
	}

	if queue.empty() != true {
		t.Errorf("Expected empty queue, was:", queue.empty())
	}

	// Test storing non-string values.
	mixed_queue := Queue{}
	mixed_queue.put(1)
	mixed_queue.put("abc")
	mixed_queue.put(3.14159)

	if mixed_queue.length() != 3 {
		t.Errorf("Expected 3 length queue, was:", mixed_queue.length())
	}

	first_mixed := mixed_queue.get()

	if first_mixed != 1 {
		t.Errorf("Expected 1, was:", first_mixed)
	}

	second_mixed := mixed_queue.get()

	if second_mixed != "abc" {
		t.Errorf("Expected \"abc\", was:", second_mixed)
	}

	third_mixed := mixed_queue.get()

	if third_mixed != 3.14159 {
		t.Errorf("Expected 3.14159, was:", third_mixed)
	}
}
