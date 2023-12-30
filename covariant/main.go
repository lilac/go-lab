package covariant

import (
	"fmt"
	"go/constant"
	"strings"
)

type Runnable interface {
	Run()
}

type BaseRunnable struct {
	Exec func()
}

func (r BaseRunnable) Run() {
	fmt.Println("Start running")
	r.Exec()
	fmt.Println("End running")
}

type Job struct {
	BaseRunnable
	name string
}

func NewJob(name string) Runnable {
	return &Job{
		BaseRunnable: BaseRunnable{
			Exec: func() {
				fmt.Printf("Running job %s\n", name)
			},
		},
		name: name,
	}
}

func join[T fmt.Stringer](elements []T) {
	ss := make([]string, len(elements))
	for i, element := range elements {
		ss[i] = element.String()
	}
	strings.Join(ss, ",")
}

func main() {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Job %d", i)
		job := NewJob(name)
		job.Run()
	}
	exps := []constant.Kind{
		constant.Int,
		constant.Bool,
	}
	join(exps)
}
