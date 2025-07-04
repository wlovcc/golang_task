package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func func_pointer01(p *int) {
	if nil == p {
		return
	}
	*p += 10
}

func func_pointer02(p []int) {
	if nil == p {
		return
	}
	for i := range p {
		p[i] *= 2
	}
}

func pointer() {
	var num = 5
	func_pointer01(&num)
	fmt.Println(num)

	nums := []int{1, 2, 3, 4, 5}
	func_pointer02(nums)
	fmt.Println(nums)
}

func func_Goroutine01() {
	var w sync.WaitGroup
	w.Add(2)

	go func() {
		defer w.Done()
		for i := 1; i <= 10; i = i + 2 {
			fmt.Println("Goroutine one: ", i)
		}
	}()

	go func() {
		defer w.Done()
		for i := 2; i <= 10; i = i + 2 {
			fmt.Println("Goroutine two: ", i)
		}
	}()
	w.Wait()
}

func task_Goroutine(i int, ch chan<- int) {
	begin := time.Now()
	duration := time.Duration(rand.Intn(2)+1) * time.Second

	for j := 0; j <= i; j++ {
		time.Sleep(duration)
	}

	res := time.Since(begin)
	fmt.Printf("task %d completed with time %s\n", i, res)
	ch <- i
}

func func_Goroutine02() {
	var ch = make(chan int, 10)
	for i := range 10 {
		go task_Goroutine(i, ch)
	}
	time_out := time.After(20 * time.Second)

	for {
		select {
		case task_num := <-ch:
			fmt.Printf("Received completion signal for task %d\n", task_num)
		case <-time_out:
			fmt.Printf("Timeout reached, exiting...")
			return
		default:
			fmt.Println("No tasks completed yet, waiting...")
			time.Sleep(500 * time.Millisecond) // Sleep to avoid busy waiting
		}
	}
}
func Goroutine() {
	func_Goroutine01()
	func_Goroutine02()
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func func_oop1() {
	var rect Rectangle
	rect.Width = 3
	rect.Height = 4

	cric := &Circle{Radius: 5}

	fmt.Println("Rectangle Area:", rect.Area())
	fmt.Println("Rectangle Perimeter:", rect.Perimeter())
	fmt.Println("Circle Area:", cric.Area())
	fmt.Println("Circle Perimeter:", cric.Perimeter())
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (emp *Employee) PrintInfo() {
	fmt.Println("Employee Name:", emp.Name, "\nAge:", emp.Age, "\nEID:", emp.EmployeeID)
}

func func_oop2() {
	emp := &Employee{
		Person:     Person{Name: "abc", Age: 23},
		EmployeeID: "115689",
	}

	emp.PrintInfo()
}

func oop() {
	func_oop1()
	func_oop2()
}

func channel() {
	w := sync.WaitGroup{}
	ch := make(chan int, 5)

	w.Add(2)

	go func() {
		defer w.Done()
		for i := range 100 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer w.Done()
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			fmt.Println("Received from channel: ", num)
		}
	}()
	w.Wait()
}

type Counter struct {
	mu    sync.Mutex
	count int
}

type Counter2 struct {
	count int64
}

func (c *Counter2) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func func_matax01() {
	var w sync.WaitGroup
	c := &Counter{}

	for range 10 {
		w.Add(1)

		go func(c *Counter, w *sync.WaitGroup) {
			defer w.Done()
			for range 1000 {
				c.mu.Lock()
				c.count++
				c.mu.Unlock()
			}
		}(c, &w)
	}
	w.Wait()

	println("Final count:", c.count)
}

func func_matax02() {
	var w sync.WaitGroup
	c2 := &Counter2{}

	for range 10 {
		w.Add(1)

		go func(c2 *Counter2, w *sync.WaitGroup) {
			defer w.Done()
			for range 1000 {
				c2.Increment()
			}
		}(c2, &w)
	}
	w.Wait()

	println("Final count:", c2.count)
}

func mutax() {
	func_matax01()
	func_matax02()
}

func main() {
	pointer()
	Goroutine()
	oop()
	channel()
	mutax()
}
