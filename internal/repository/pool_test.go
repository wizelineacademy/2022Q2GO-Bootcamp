package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testJob struct {
	id   int
	task interface{}
}

func (tk testJob) getId() int {
	return tk.id
}

func (tk testJob) getTask() interface{} {
	return tk.task
}

func (tk testJob) RunTask() interface{} {
	time.Sleep(time.Second)

	fmt.Printf("Finished task #%v\n", tk.id)

	return fmt.Sprintf("task #%v\n", tk.id)
}

func TestDispatcher(t *testing.T) {
	pool := NewGoroutinePool(5)
	taskSize := 10

	var tasks []Job
	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, testJob{
			id:   v,
			task: fmt.Sprintf("%d", v),
		})
	}

	go pool.AllocateJobs(tasks)

	done := make(chan bool)

	go pool.GetResult(done)

	pool.AddWorkers()
	<-done

	assert.NotNil(t, pool)
	assert.EqualValues(t, len(pool.ResultsPool), taskSize)
}
