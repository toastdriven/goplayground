package queue

import (
	"math"
	"sync"
)

type Queue struct {
	data        []interface{}
	current_min int
	lock        sync.Mutex
}

func (q *Queue) put(val interface{}) {
	q.lock.Lock()
	current_length := q.length()
	current_cap := cap(q.data)

	if current_length+1 >= current_cap {
		new_cap := int(math.Floor(float64(current_cap) * 1.25))
		new_data := make([]interface{}, current_length, new_cap+1)
		copy(new_data, q.data)
		q.data = new_data
	}

	q.data = append(q.data, val)
	q.lock.Unlock()
}

func (q *Queue) get() interface{} {
	q.lock.Lock()
	val := q.data[q.current_min]
	q.current_min++

	// To prevent lots of ``make`` churn, we only resize every 10 items.
	if q.current_min >= 10 {
		copy(q.data, q.data[q.current_min:])
		q.current_min = 0
	}

	q.lock.Unlock()
	return val
}

func (q *Queue) length() int {
	return len(q.data) - q.current_min
}

func (q *Queue) empty() bool {
	return q.length() <= 0
}
