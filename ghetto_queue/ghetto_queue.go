package main

import (
	"fmt"
	"math"
)

type Queue struct {
	data        []string
	current_min int
}

func (q *Queue) put(val string) {
	current_length := q.length()
	current_cap := cap(q.data)

	if current_length+1 >= current_cap {
		new_cap := int(math.Floor(float64(current_cap) * 1.25))
		new_data := make([]string, current_length, new_cap+1)
		copy(new_data, q.data)
		q.data = new_data
	}

	q.data = append(q.data, val)
}

func (q *Queue) get() string {
	val := q.data[q.current_min]
	q.current_min++

	// To prevent lots of ``make`` churn, we only resize every 10 items.
	if q.current_min >= 10 {
		copy(q.data, q.data[q.current_min:])
		q.current_min = 0
	}

	return val
}

func (q *Queue) length() int {
	return len(q.data) - q.current_min
}

func (q *Queue) empty() bool {
	return q.length() <= 0
}

func main() {
	// Just for testing.
	queue := Queue{}
	fmt.Println("Len:", queue.length())
	queue.put("Hello")
	queue.put("World")
	queue.put("foo")
	fmt.Println("Len:", queue.length())
	fmt.Println("Empty?", queue.empty())
	first := queue.get()
	fmt.Println("Saw:", first)
	fmt.Println("Len:", queue.length())
	second := queue.get()
	fmt.Println("Saw:", second)
	fmt.Println("Empty?", queue.empty())
	third := queue.get()
	fmt.Println("Saw:", third)
	fmt.Println("Len:", queue.length())
	fmt.Println("Empty?", queue.empty())
}
