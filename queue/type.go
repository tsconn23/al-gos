package queue

type Queue struct {}

func(q *Queue) Enqueue() string {
	return "enqueued"
}

func(q *Queue) Dequeue() string {
	return "dequeued"
}
