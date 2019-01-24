package jobs_test

import (
	"sync"
	"testing"

	"github.com/joincivil/go-common/pkg/jobs"
)

type Spy struct {
	RunCount int
}

func NewSpy() *Spy {
	return &Spy{
		RunCount: 0,
	}
}

func (s *Spy) Run() {
	s.RunCount = s.RunCount + 1
}

func buildJob(spy *Spy) *jobs.Job {
	jobID := "test"
	work := func(updates chan<- string) {
		updates <- "foo"
		updates <- "bar"
		spy.Run()
	}
	return jobs.NewJob(jobID, work)
}

func TestJob(t *testing.T) {

	t.Run("run work", func(t *testing.T) {
		spy := NewSpy()
		job := buildJob(spy)
		if job.GetStatus() != "initialized" {
			t.Fatalf("job status should be `initialized`")
		}
		if spy.RunCount > 0 {
			t.Fatalf("work function should not have run")
		}

		job.Start()
		job.WaitForFinish()
		if job.GetStatus() != "complete" {
			t.Fatalf("job status should be `complete`")
		}
		if spy.RunCount != 1 {
			t.Fatalf("work function should be 1")
		}
	})

	t.Run("subscriptions", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		spy := NewSpy()
		spySub1 := NewSpy()
		spySub2 := NewSpy()
		job := buildJob(spy)
		sub1 := job.Subscribe()
		sub2 := job.Subscribe()
		job.Start()

		go func() {
			for range sub1.Updates {
				spySub1.Run()
			}
			wg.Done()
		}()

		go func() {
			for range sub2.Updates {
				spySub2.Run()
			}
			wg.Done()
		}()

		wg.Wait()

		if job.GetStatus() != "complete" {
			t.Fatalf("job status should be `complete`")
		}
		if spy.RunCount != 1 {
			t.Fatalf("work RunCount should be 1 but is %v", spy.RunCount)
		}
		if spySub1.RunCount != 2 {
			t.Fatalf("work RunCount should be 2 but is %v", spySub1.RunCount)
		}
		if spySub2.RunCount != 2 {
			t.Fatalf("work RunCount should be 2 but is %v", spySub2.RunCount)
		}
	})

	t.Run("unsubscribe", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)

		spyJob := NewSpy()
		spySub1 := NewSpy()
		spySub2 := NewSpy()

		job := buildJob(spyJob)
		sub1 := job.Subscribe()
		sub2 := job.Subscribe()

		job.Unsubscribe(sub2)
		job.Start()

		go func() {
			for range sub1.Updates {
				spySub1.Run()
			}
			wg.Done()
		}()

		go func() {
			for range sub2.Updates {
				spySub2.Run()
			}
			wg.Done()
		}()

		wg.Wait()

		if spyJob.RunCount != 1 {
			t.Fatalf("work RunCount should be 1 but is %v", spyJob.RunCount)
		}
		if spySub1.RunCount != 2 {
			t.Fatalf("spySub3 RunCount should be 0 but is %v", spyJob.RunCount)
		}
		if spySub2.RunCount != 0 {
			t.Fatalf("spySub3 RunCount should be 0 but is %v", spySub2.RunCount)
		}

	})
}
