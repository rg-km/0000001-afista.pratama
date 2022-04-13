package main

import "fmt"

// struct + array / slice
func NewQueue(data []int) *queue {
	return &queue{
		datas: data,
	}
}

type queue struct {
	datas []int
}

func (q *queue) Enqueue(newData int) {
	fmt.Println("Enqueue: add data ", newData)
	q.datas = append(q.datas, newData)
	fmt.Println(q.datas)
}

func (q *queue) Dequeue() {
	fmt.Println("Dequeue: prosess data", q.datas[0])
	q.datas = q.datas[1:]
	fmt.Println(q.datas)
}

func main() {
	queue := NewQueue([]int{1, 2, 3, 4, 5, 6, 7})

	queue.Enqueue(8)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	// queue := []int{1, 2, 3, 4, 5, 6}

	// q1 := Enqueue(queue, 7)
	// fmt.Println(q1)

	// d1 := Dequeue(q1)
	// fmt.Println(d1)

	// d2 := Dequeue(d1)
	// fmt.Println(d2)

}

func Enqueue(arr []int, newData int) []int {
	fmt.Println("Enqueue: ", newData)

	return append(arr, newData)
}

func Dequeue(arr []int) []int {
	fmt.Println("Dequeue: ", arr[0])

	return arr[1:]
}

// karakteristik dari queue
// FIFO, hanya bisa proses dari depan
// menambah di belakang
// tidak bisa merusak urutan / antrian (private data)
