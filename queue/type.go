package queue

type Queue struct {}

func(q *Queue) enqueue() string {
	return "enqueued"
}

func(q *Queue) dequeue() string {
	return "dequeued"
}
