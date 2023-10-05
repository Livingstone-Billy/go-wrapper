package main

import (
	"errors"
)

// implementing queue
type Queue struct {
	data []interface{}
}

func (q *Queue) push(item interface{}) {
	q.data = append(q.data, item)
}

func (q *Queue) top() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := q.data[0]
	return item, nil
}

func (q *Queue) pop() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("Queue is empty")
	}
	firstItem := q.data[0]
	q.data = q.data[1:]
	return firstItem, nil
}

func (q *Queue) size() int {
	return len(q.data)
}

func (q *Queue) isEmpty() bool {
	if len(q.data) == 0 {
		return true
	} else {
		return false
	}
}
